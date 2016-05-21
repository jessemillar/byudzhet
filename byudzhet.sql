-- phpMyAdmin SQL Dump
-- version 4.0.10deb1
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: May 20, 2016 at 10:04 AM
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
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=16 ;

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
  KEY `user` (`user`),
  KEY `bucket` (`bucket`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=33 ;

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
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 AUTO_INCREMENT=5 ;

-- --------------------------------------------------------

--
-- Table structure for table `projected`
--

CREATE TABLE IF NOT EXISTS `projected` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `user` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  PRIMARY KEY (`ID`)
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

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
