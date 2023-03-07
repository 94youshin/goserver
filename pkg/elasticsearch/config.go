package elasticsearch

type Configuration struct {
	Servers []string
}

func (c *Configuration) NewClient() {

}
