package thunderbirdparser

import "net/http"
import "encoding/xml"
import "sync"
import "regexp"
import "log"
import "time"

type ThunderbirdParser struct {
	httpClient *http.Client
	serviceUrl string 
	threads int
	domainSettings map[string]ClientConfig
}

func ThunderbirdParserCreate(threads int) (p *ThunderbirdParser) {
	p = new(ThunderbirdParser)
	p.serviceUrl = "https://autoconfig.thunderbird.net/v1.1/"
	p.httpClient = new(http.Client)
	p.threads = threads
	p.domainSettings = make(map[string]ClientConfig)
	return
}

func (p *ThunderbirdParser) ParseUrl(docUrl string) (config ClientConfig, err error) {
	data, err := httpGet(docUrl, p.httpClient)
	if err != nil {
		return
	}
	err = xml.Unmarshal(data, &config)
	return
}

func (p *ThunderbirdParser) DownloadAll() (domainSettings map[string]ClientConfig, err error) {
	data, err := httpGet(p.serviceUrl, p.httpClient)
	if err != nil {
		return
	}
	
	exp := regexp.MustCompile("> <a href=\"(.+?)\">")
	configUrls := exp.FindAllStringSubmatch(string(data), -1)

	configUrlChan := make(chan string)
	configChan := make(chan ClientConfig)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(p.threads)

	for i:=0; i < p.threads; i++ {
		go func() {
			defer func() {
				waitGroup.Done()
			}()
			for {
				select {
					case configUrl := <-configUrlChan:
						if configUrl == "" {
							return
						}
						log.Printf("parse url: %s", configUrl)
						config, err := p.ParseUrl(configUrl)
						if err == nil {
							configChan <- config	
						}
					case <- time.After(time.Second):
						return
				}
			}
		}()
	}

	go func() {
		defer waitGroup.Done()
		for config := range configChan {
			for _, ep := range config.EmailProviders {
				for _, domain := range ep.Domains {
					p.domainSettings[domain] = config
				}
			}
		}
	}()

	for _, match := range configUrls {
		configUrlChan <- p.serviceUrl + match[1]
	}
	close(configUrlChan)
	waitGroup.Wait()
	domainSettings = p.domainSettings
	return
}

func (p *ThunderbirdParser) SearchDomain(domain string) ClientConfig {
	return p.domainSettings[domain]
}