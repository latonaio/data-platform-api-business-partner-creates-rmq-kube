package existence_conf

import "data-platform-api-business-partner-creates-rmq-kube/config"

type exconfQueueMapper map[string]string

func NewExconfQueueMapper(c *config.Conf) exconfQueueMapper {
	m := exconfQueueMapper{}
	ecQNames := c.RMQ.QueueToExConf()
	m["BusinessPartnerGeneral"] = ecQNames[0%len(ecQNames)]
	m["Industry"] = ecQNames[1%len(ecQNames)]
	m["Country"] = ecQNames[2%len(ecQNames)]
	m["Language"] = ecQNames[3%len(ecQNames)]
	m["Currency"] = ecQNames[4%len(ecQNames)]
	return m
}

func (m exconfQueueMapper) getQueueNameByConfContent(content string) string {
	return m[content]
}
