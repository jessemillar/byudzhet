-- phpMyAdmin SQL Dump
-- version 4.0.10deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: May 10, 2016 at 01:54 PM
-- Server version: 5.5.49-0ubuntu0.14.04.1
-- PHP Version: 5.5.9-1ubuntu4.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `byudzhet`
--

-- --------------------------------------------------------

--
-- Table structure for table `buckets`
--

CREATE TABLE IF NOT EXISTS `buckets` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `user` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `name` varchar(140) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `user` (`user`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=2 ;

--
-- Dumping data for table `buckets`
--

INSERT INTO `buckets` (`ID`, `user`, `amount`, `name`) VALUES
(1, 8, 100, 'Poots');

-- --------------------------------------------------------

--
-- Table structure for table `expenses`
--

CREATE TABLE IF NOT EXISTS `expenses` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `user` int(11) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `bucket` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `recipient` varchar(50) NOT NULL,
  `note` text NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `user` (`user`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=19 ;

--
-- Dumping data for table `expenses`
--

INSERT INTO `expenses` (`ID`, `user`, `time`, `bucket`, `amount`, `recipient`, `note`) VALUES
(15, 8, '2016-05-10 00:47:23', 0, 350, 'Target', 'PS4'),
(16, 8, '2016-04-10 00:47:32', 0, 50, 'Target', 'Uncharted 4'),
(17, 8, '2016-05-10 02:12:54', 1, 10, 'Smith''s', 'Antacid'),
(18, 10, '2016-05-10 14:56:45', 0, 200, 'Stuff', 'Poots');

-- --------------------------------------------------------

--
-- Table structure for table `income`
--

CREATE TABLE IF NOT EXISTS `income` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `user` int(11) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `payer` varchar(140) NOT NULL,
  `amount` int(11) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `user` (`user`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=2 ;

--
-- Dumping data for table `income`
--

INSERT INTO `income` (`ID`, `user`, `time`, `payer`, `amount`) VALUES
(1, 8, '2016-05-10 17:00:32', 'BYU', 500);

-- --------------------------------------------------------

--
-- Table structure for table `sharing`
--

CREATE TABLE IF NOT EXISTS `sharing` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `user` int(11) NOT NULL,
  `sharee` int(11) NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `user` (`user`,`sharee`),
  KEY `receiver` (`sharee`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 AUTO_INCREMENT=1 ;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE IF NOT EXISTS `users` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(140) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=11 ;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`ID`, `email`) VALUES
(8, 'hellojessemillar@gmail.com'),
(10, '');

--
-- Constraints for dumped tables
--

--
-- Constraints for table `buckets`
--
ALTER TABLE `buckets`
  ADD CONSTRAINT `buckets_ibfk_1` FOREIGN KEY (`user`) REFERENCES `users` (`ID`);

--
-- Constraints for table `expenses`
--
ALTER TABLE `expenses`
  ADD CONSTRAINT `expenses_ibfk_1` FOREIGN KEY (`user`) REFERENCES `users` (`ID`);

--
-- Constraints for table `income`
--
ALTER TABLE `income`
  ADD CONSTRAINT `income_ibfk_1` FOREIGN KEY (`user`) REFERENCES `users` (`ID`);

--
-- Constraints for table `sharing`
--
ALTER TABLE `sharing`
  ADD CONSTRAINT `sharing_ibfk_2` FOREIGN KEY (`sharee`) REFERENCES `users` (`ID`),
  ADD CONSTRAINT `sharing_ibfk_1` FOREIGN KEY (`user`) REFERENCES `users` (`ID`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
