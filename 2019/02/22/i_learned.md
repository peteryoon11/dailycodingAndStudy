## spring boot batch job config


/**
 * 
 */
package com.example.demo.config;


import javax.sql.DataSource;
import java.sql.SQLException;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.EnableAspectJAutoProxy;
import org.springframework.transaction.annotation.EnableTransactionManagement;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;
import org.springframework.context.annotation.Primary;

import org.springframework.context.annotation.Bean;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.boot.context.properties.ConfigurationProperties;


import com.zaxxer.hikari.HikariDataSource;


/**
 * @author js.yoon
 *
 */
@Configuration
@EnableTransactionManagement
@EnableJpaRepositories
public class BatchJobConfig {
	@Primary
	@Bean(name = "jobadmDataSource")
	@Qualifier("jobadmDataSource")
	@ConfigurationProperties(prefix = "datasource.jobadm")
	public DataSource jobadmDataSource() throws SQLException{
		HikariDataSource jobadmDataSource = new HikariDataSource();
		return jobadmDataSource;
	}
}

