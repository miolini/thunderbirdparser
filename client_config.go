package thunderbirdparser

type ClientConfig struct {
	EmailProviders []EmailProvider `xml:"emailProvider"`
	Documentation []Documentation `xml:"documentation"`
}

type EmailProvider struct {
	Id string `xml:"id,attr"`
	Domains []string `xml:"domain"`
	DisplayName string `xml:"displayName"`
	DisplayShortName string `xml:"displayShortName"`
	IncomingServers []IncomingServer `xml:"incomingServer"`
	OutgoingServers []OutgoingServer `xml:"outgoingServer"`
	Enable Enable `xml:"enable"`
}

type IncomingServer struct {
	Type string `xml:"type,attr"`
	Hostname string `xml:"hostname"`
	Port int `xml:"port"`
	SocketType string `xml:"socketType"`
	Username string `xml:"username"`
	Authentication string `xml:"authentication"`
}

type OutgoingServer struct {
	Type string `xml:"type,attr"`
	Hostname string `xml:"hostname"`
	Port int `xml:"port"`
	SocketType string `xml:"socketType"`
	Username string `xml:"username"`
	Authentication string `xml:"authentication"`
}

type Enable struct {
	VisitUrl string `xml:"visiturl,attr"`
	Instruction string `xml:"instruction"`
}

type Documentation struct {
	Url string `xml:"url,attr"`
	Descr string `xml:"descr"`
}