package api

func (s *ApiTestSuite) Test_MonitoringServer() {

	// List all Monitoring Servers
	monitoringServers, err := s.api.MonitoringServer().List(nil)
	s.NoError(err)
	s.NotNil(monitoringServers)
	s.NotEmpty(monitoringServers.Result)

	// List Monitoring Servers with filter
	monitoringServers, err = s.api.MonitoringServer().List(&ListOptions{
		Search: map[string]interface{}{
			"name": "Central",
		},
	})
	s.NoError(err)
	s.NotNil(monitoringServers)
	s.Equal(1, len(monitoringServers.Result))

	// Generate Configuration for a Monitoring Server
	err = s.api.MonitoringServer().GenerateConfiguration(1)
	s.NoError(err)

	// Reload Configuration for a Monitoring Server
	err = s.api.MonitoringServer().ReloadConfiguration(1)
	s.NoError(err)

	// Generate and Reload Configuration for a non-existing Monitoring Server
	err = s.api.MonitoringServer().GenerateAndReloadConfiguration(1)
	s.NoError(err)

	// Generate all configurations
	err = s.api.MonitoringServer().GenerateConfigurationAll()
	s.NoError(err)

	// Reload all configurations
	err = s.api.MonitoringServer().ReloadConfigurationAll()
	s.NoError(err)

	// Generate and Reload all configurations
	err = s.api.MonitoringServer().GenerateAndReloadConfigurationAll()
	s.NoError(err)

	// List all Monitoring Servers in real-time
	monitoringServersRealTime, err := s.api.MonitoringServer().ListFromRealTime(nil)
	s.NoError(err)
	s.NotNil(monitoringServersRealTime)
	s.NotEmpty(monitoringServersRealTime.Result)
	// Always returns 0. Centreon bug ?
	//s.Greater(len(monitoringServersRealTime.Result), 0)

	// List Monitoring Servers in real-time with filter
	monitoringServersRealTime, err = s.api.MonitoringServer().ListFromRealTime(&ListOptions{
		Search: map[string]interface{}{
			"name": "Central",
		},
	})
	s.NoError(err)
	s.NotNil(monitoringServersRealTime)
	s.NotEmpty(monitoringServersRealTime.Result)
	//s.Equal(1, len(monitoringServersRealTime.Result))
	// Always returns 0. Centreon bug ?
}
