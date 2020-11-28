CREATE TABLE `Malware` (
    `id` int  NOT NULL ,
    `malware_type_id` int   ,
    `situation` varchar(20)   ,
    PRIMARY KEY (
        `id`
    )
);

CREATE TABLE `MalwareType` (
    `malware_type_id` int  NOT NULL ,
    `malware_type` varchar(20)   ,
    `malware_version` varchar(30)   ,
    `malware_release_date` datetime   ,
    `target_system` varchar(20)   ,
    `atack_vector_id` int   ,
    PRIMARY KEY (
        `malware_type_id`
    )
);

CREATE TABLE `AttackVector` (
    `attack_vector_id` int  NOT NULL ,
    `attack_vector_type` varchar(20)   ,
    `embedded_file` varchar   ,
    PRIMARY KEY (
        `attack_vector_id`
    )
);

CREATE TABLE `MalwareStatus` (
    `id` int   ,
    `create_date` datetime   ,
    `infected_date` datetime   ,
    `first_touch_with_cc` datetime   ,
    `clean_date` datetime   
);

CREATE TABLE `Money` (
    `malware_id` int   ,
    `amount` float8   ,
    `status` int   
);

CREATE TABLE `Accounts` (
    `account_id` int  NOT NULL ,
    `product` varchar(20)   ,
    `username` varchar(20)   ,
    `password` varchar(20)   ,
    `email` varchar(50)   ,
    `create_date` datetime   ,
    PRIMARY KEY (
        `account_id`
    )
);

CREATE TABLE `MiddleWareIP` (
    `id` int  NOT NULL ,
    `ip` varchar(20)   ,
    `account_id` int   ,
    PRIMARY KEY (
        `id`
    )
);

CREATE TABLE `Victim` (
    `id` int   ,
    `victim_ip` varchar(16)   ,
    `victim_local_ip` varchar(16)   ,
    `computer_name` varchar(25)   ,
    `username` varchar(28)   ,
    `computer_ram` float   ,
    `computer_cpu` varchar(35)   ,
    `computer_status` varbinary   ,
    `botnet_status` varbinary   
);

CREATE TABLE `Botnet` (
    `victim_id` int   ,
    `port` int   ,
    `protocol` int   ,
    `token` varchar(45)   
);

CREATE TABLE `MiddleWareDomain` (
    `id` int  NOT NULL ,
    `domain_name` varchar(50)  ,
    `status` int   ,
    `which_domain_name_provider` varchar(25)   ,
    `account_id` int   ,
    PRIMARY KEY (
        `id`
    )
);

CREATE TABLE 'IPWhois' (
    'ip' varchar(16) NOT NULL ,
    'isp' varchar(20),
    'country' varchar(15),
    'city' varchar(30),
    'longtitude' float8,
    'latitude' float8,
    PRIMARY KEY (
    	'ip'
    )
);



ALTER TABLE `Malware` ADD CONSTRAINT `fk_Malware_malware_type_id` FOREIGN KEY(`malware_type_id`)
REFERENCES `MalwareType` (`malware_type_id`);

ALTER TABLE `MalwareType` ADD CONSTRAINT `fk_MalwareType_atack_vector_id` FOREIGN KEY(`atack_vector_id`)
REFERENCES `AttackVector` (`attack_vector_id`);

ALTER TABLE `MalwareStatus` ADD CONSTRAINT `fk_MalwareStatus_id` FOREIGN KEY(`id`)
REFERENCES `Malware` (`id`);

ALTER TABLE `Money` ADD CONSTRAINT `fk_Money_malware_id` FOREIGN KEY(`malware_id`)
REFERENCES `Malware` (`id`);

ALTER TABLE `MiddleWareIP` ADD CONSTRAINT `fk_MiddleWareIP_id` FOREIGN KEY(`id`)
REFERENCES `IPWhois` (`ip`);

ALTER TABLE `MiddleWareIP` ADD CONSTRAINT `fk_MiddleWareIP_account_id` FOREIGN KEY(`account_id`)
REFERENCES `Accounts` (`account_id`);

ALTER TABLE `Victim` ADD CONSTRAINT `fk_Victim_id` FOREIGN KEY(`id`)
REFERENCES `Malware` (`id`);

ALTER TABLE `Victim` ADD CONSTRAINT `fk_Victim_victim_ip` FOREIGN KEY(`victim_ip`)
REFERENCES `MiddleWareIP` (`ip`);

ALTER TABLE `Botnet` ADD CONSTRAINT `fk_Botnet_victim_id` FOREIGN KEY(`victim_id`)
REFERENCES `Victim` (`id`);

ALTER TABLE `MiddleWareDomain` ADD CONSTRAINT `fk_MiddleWareDomain_account_id` FOREIGN KEY(`account_id`)
REFERENCES `Accounts` (`account_id`);


