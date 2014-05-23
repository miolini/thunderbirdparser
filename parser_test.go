package thunderbirdparser

import "testing"
// import "encoding/json"
// import "encoding/xml"

func TestDownloadAll(t *testing.T) {
	var err error
	p := ThunderbirdParserCreate(10)
	t.Logf("test thunderbird parser: download all")
	settings, err := p.DownloadAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("founded %d domains", len(settings))
}

/*
func TestDownloadNarodRu(t *testing.T) {
	p := ThunderbirdParserCreate(4)
	t.Logf("test thunderbird parser")
	clientConfig, err := p.ParseUrl("https://autoconfig.thunderbird.net/v1.1/narod.ru")
	if err != nil {
		t.Fatal(err)
	}
	dump, err := json.Marshal(clientConfig)
	t.Logf("clientConfig = %s", dump)
}

func TestDownloadTjekketDk(t *testing.T) {
	p := ThunderbirdParserCreate(4)
	t.Logf("test thunderbird parser")
	clientConfig, err := p.ParseUrl("https://autoconfig.thunderbird.net/v1.1/tjekket.dk")
	if err != nil {
		t.Fatal(err)
	}
	dump, err := xml.MarshalIndent(clientConfig, "", "   ")
	t.Logf("clientConfig = %s", dump)	
}
*/