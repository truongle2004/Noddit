package com.config_service.config.service;

import org.springframework.boot.SpringApplication;
import org.springframework.cloud.config.server.EnableConfigServer;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@EnableConfigServer
public class NodditApplication {

	public static void main(String[] args) {
		SpringApplication.run(NodditApplication.class, args);
	}

}
