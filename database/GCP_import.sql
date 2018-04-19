CREATE DATABASE `lab_inventory`;
USE `lab_inventory`;

--
-- Table structure for table `config_table`
--

DROP TABLE IF EXISTS `config_table`;
CREATE TABLE `config_table` (
  `config_file` varchar(255) DEFAULT NULL,
  `config_path` varchar(255) DEFAULT NULL
);

--
-- Dumping data for table `config_table`
--

LOCK TABLES `config_table` WRITE;
INSERT INTO `config_table` VALUES ('nleiva-default','/config/nleiva-default/mrstn-1002-1.cisco.com'),
('milally-even-exr','/config/milally-evpn-exr/mrstn-5501-3.cisco.com');
UNLOCK TABLES;

--
-- Table structure for table `device_table`
--

DROP TABLE IF EXISTS `device_table`;
CREATE TABLE `device_table` (
  `node` varchar(255) DEFAULT NULL,
  `sw_image` varchar(255) DEFAULT NULL,
  `hardware` varchar(255) DEFAULT NULL,
  `config` varchar(255) DEFAULT NULL,
  `user` varchar(255) DEFAULT NULL,
  `device_status` varchar(255) DEFAULT NULL
);

--
-- Dumping data for table `device_table`
--

LOCK TABLES `device_table` WRITE;
INSERT INTO `device_table` VALUES ('mrstn-1002-1','6.2.25','ASR-1002','access','none','offline'),
('mrstn-1002-2','6.5.1','ASR-1002','access','seyost','online'),
('mrstn-5501-1','6.5.2','NCS-5501-SE','leaf','nkumar','online'),
('mrstn-5501-2','6.2.1','NCS-5501-SE','leaf','nkumar','offline'),
('mrstn-5501-3','6.3.1','NCS-5501-SE','leaf','nleiva','online'),
('mrstn-5501-4','6.4.1','NCS-5501-SE','leaf','nleiva','online'),
('mrstn-5502-1','6.3.2','NCS-5502-SE','spine','seyost','online'),
('mrstn-5502-2','6.5.6','NCS-5502-SE','spine','nkumar','offline'),
('mrstn-5502-3','6.5.17','NCS-5502-SE','spine','nleiva','online');
UNLOCK TABLES;

--
-- Table structure for table `software_table`
--

DROP TABLE IF EXISTS `software_table`;
CREATE TABLE `software_table` (
  `Filename` varchar(255) DEFAULT NULL,
  `Imagename` varchar(255) DEFAULT NULL
);

--
-- Dumping data for table `software_table`
--

LOCK TABLES `software_table` WRITE;
INSERT INTO `software_table` VALUES ('ncs-6.5.1.iso','6.5.1'),
('ASR-6.2.25.iso','6.2.25'),
('ncs-6.3.2.iso','6.3.2');
UNLOCK TABLES;

--
-- Table structure for table `user_table`
--

DROP TABLE IF EXISTS `user_table`;
CREATE TABLE `user_table` (
  `name` varchar(255) DEFAULT NULL,
  `Type` varchar(255) DEFAULT NULL
);

--
-- Dumping data for table `user_table`
--

LOCK TABLES `user_table` WRITE;
INSERT INTO `user_table` VALUES ('seyost','GOD'),
('nleiva','admin'),
('nkumar','admin');
UNLOCK TABLES;