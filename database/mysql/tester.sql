-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: Nov 28, 2020 at 11:46 AM
-- Server version: 5.7.32
-- PHP Version: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tester`
--

-- --------------------------------------------------------

--
-- Table structure for table `Accounts`
--

CREATE TABLE `Accounts` (
  `account_id` int(11) NOT NULL,
  `product` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `username` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `password` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `email` varchar(50) COLLATE utf8_turkish_ci DEFAULT NULL,
  `create_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `AttackVector`
--

CREATE TABLE `AttackVector` (
  `attack_vector_id` int(11) NOT NULL,
  `attack_vector_type` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `embedded_file` varchar(30) COLLATE utf8_turkish_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Botnet`
--

CREATE TABLE `Botnet` (
  `victim_id` int(11) DEFAULT NULL,
  `port` int(11) DEFAULT NULL,
  `protocol` int(11) DEFAULT NULL,
  `token` varchar(45) COLLATE utf8_turkish_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `IPWhois`
--

CREATE TABLE `IPWhois` (
  `ip` varchar(20) COLLATE utf8_turkish_ci NOT NULL,
  `isp` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `country` varchar(15) COLLATE utf8_turkish_ci DEFAULT NULL,
  `city` varchar(30) COLLATE utf8_turkish_ci DEFAULT NULL,
  `longtitude` double DEFAULT NULL,
  `latitude` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Malware`
--

CREATE TABLE `Malware` (
  `id` int(11) NOT NULL,
  `malware_type_id` int(11) DEFAULT NULL,
  `situation` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `MalwareStatus`
--

CREATE TABLE `MalwareStatus` (
  `id` int(11) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `infected_date` datetime DEFAULT NULL,
  `first_touch_with_cc` datetime DEFAULT NULL,
  `clean_date` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `MalwareType`
--

CREATE TABLE `MalwareType` (
  `malware_type_id` int(11) NOT NULL,
  `malware_type` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `malware_version` varchar(30) COLLATE utf8_turkish_ci DEFAULT NULL,
  `malware_release_date` datetime DEFAULT NULL,
  `target_system` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `atack_vector_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `MiddleWareDomain`
--

CREATE TABLE `MiddleWareDomain` (
  `id` int(11) NOT NULL,
  `domain_name` varchar(50) COLLATE utf8_turkish_ci DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `which_domain_name_provider` varchar(25) COLLATE utf8_turkish_ci DEFAULT NULL,
  `account_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `MiddleWareIP`
--

CREATE TABLE `MiddleWareIP` (
  `id` int(11) NOT NULL,
  `ip` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `account_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Money`
--

CREATE TABLE `Money` (
  `malware_id` int(11) DEFAULT NULL,
  `amount` double DEFAULT NULL,
  `status` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Victim`
--

CREATE TABLE `Victim` (
  `id` int(11) DEFAULT NULL,
  `victim_ip` varchar(20) COLLATE utf8_turkish_ci DEFAULT NULL,
  `victim_local_ip` varchar(16) COLLATE utf8_turkish_ci DEFAULT NULL,
  `computer_name` varchar(25) COLLATE utf8_turkish_ci DEFAULT NULL,
  `username` varchar(28) COLLATE utf8_turkish_ci DEFAULT NULL,
  `computer_ram` float DEFAULT NULL,
  `computer_cpu` varchar(35) COLLATE utf8_turkish_ci DEFAULT NULL,
  `computer_status` varbinary(2) DEFAULT NULL,
  `botnet_status` varbinary(2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_turkish_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Accounts`
--
ALTER TABLE `Accounts`
  ADD PRIMARY KEY (`account_id`);

--
-- Indexes for table `AttackVector`
--
ALTER TABLE `AttackVector`
  ADD PRIMARY KEY (`attack_vector_id`);

--
-- Indexes for table `Botnet`
--
ALTER TABLE `Botnet`
  ADD KEY `fk_Botnet_victim_id` (`victim_id`);

--
-- Indexes for table `IPWhois`
--
ALTER TABLE `IPWhois`
  ADD PRIMARY KEY (`ip`);

--
-- Indexes for table `Malware`
--
ALTER TABLE `Malware`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_Malware_malware_type_id` (`malware_type_id`);

--
-- Indexes for table `MalwareStatus`
--
ALTER TABLE `MalwareStatus`
  ADD KEY `fk_MalwareStatus_id` (`id`);

--
-- Indexes for table `MalwareType`
--
ALTER TABLE `MalwareType`
  ADD PRIMARY KEY (`malware_type_id`),
  ADD KEY `fk_MalwareType_atack_vector_id` (`atack_vector_id`);

--
-- Indexes for table `MiddleWareDomain`
--
ALTER TABLE `MiddleWareDomain`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_MiddleWareDomain_account_id` (`account_id`);

--
-- Indexes for table `MiddleWareIP`
--
ALTER TABLE `MiddleWareIP`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_MiddleWareIP_account_id` (`account_id`);

--
-- Indexes for table `Money`
--
ALTER TABLE `Money`
  ADD KEY `fk_Money_malware_id` (`malware_id`);

--
-- Indexes for table `Victim`
--
ALTER TABLE `Victim`
  ADD KEY `fk_Victim_id` (`id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Botnet`
--
ALTER TABLE `Botnet`
  ADD CONSTRAINT `fk_Botnet_victim_id` FOREIGN KEY (`victim_id`) REFERENCES `Victim` (`id`);

--
-- Constraints for table `Malware`
--
ALTER TABLE `Malware`
  ADD CONSTRAINT `fk_Malware_malware_type_id` FOREIGN KEY (`malware_type_id`) REFERENCES `MalwareType` (`malware_type_id`);

--
-- Constraints for table `MalwareStatus`
--
ALTER TABLE `MalwareStatus`
  ADD CONSTRAINT `fk_MalwareStatus_id` FOREIGN KEY (`id`) REFERENCES `Malware` (`id`);

--
-- Constraints for table `MalwareType`
--
ALTER TABLE `MalwareType`
  ADD CONSTRAINT `fk_MalwareType_atack_vector_id` FOREIGN KEY (`atack_vector_id`) REFERENCES `AttackVector` (`attack_vector_id`);

--
-- Constraints for table `MiddleWareDomain`
--
ALTER TABLE `MiddleWareDomain`
  ADD CONSTRAINT `fk_MiddleWareDomain_account_id` FOREIGN KEY (`account_id`) REFERENCES `Accounts` (`account_id`);

--
-- Constraints for table `MiddleWareIP`
--
ALTER TABLE `MiddleWareIP`
  ADD CONSTRAINT `fk_MiddleWareIP_account_id` FOREIGN KEY (`account_id`) REFERENCES `Accounts` (`account_id`);

--
-- Constraints for table `Money`
--
ALTER TABLE `Money`
  ADD CONSTRAINT `fk_Money_malware_id` FOREIGN KEY (`malware_id`) REFERENCES `Malware` (`id`);

--
-- Constraints for table `Victim`
--
ALTER TABLE `Victim`
  ADD CONSTRAINT `fk_Victim_id` FOREIGN KEY (`id`) REFERENCES `Malware` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
