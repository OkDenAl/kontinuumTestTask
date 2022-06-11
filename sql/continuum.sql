-- phpMyAdmin SQL Dump
-- version 5.1.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jun 11, 2022 at 01:45 PM
-- Server version: 5.7.24
-- PHP Version: 8.0.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `continuum`
--

-- --------------------------------------------------------

--
-- Table structure for table `data_about_lessons`
--

CREATE TABLE `data_about_lessons` (
  `lesson_date` date DEFAULT NULL,
  `idGroup` int(11) DEFAULT NULL,
  `idStudent` int(11) DEFAULT NULL,
  `activity` int(11) DEFAULT NULL,
  `HW` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `data_about_lessons`
--

INSERT INTO `data_about_lessons` (`lesson_date`, `idGroup`, `idStudent`, `activity`, `HW`) VALUES
('2021-09-27', 1, 1, 5, 100),
('2021-09-27', 1, 2, 3, 91),
('2021-09-27', 1, 4, 3, 45),
('2021-09-30', 2, 5, 3, 60),
('2021-09-30', 2, 2, 1, 25),
('2021-09-30', 2, 3, 1, 35),
('2021-09-30', 2, 1, 4, 10),
('2021-10-04', 1, 1, 2, 73),
('2021-10-04', 1, 2, 4, 58),
('2021-10-04', 1, 4, 3, 82),
('2021-10-07', 2, 5, 5, 80),
('2021-10-07', 2, 2, 5, 90),
('2021-10-07', 2, 3, 5, 100),
('2021-10-07', 2, 1, 4, 95),
('2021-10-11', 1, 1, 4, 45),
('2021-10-11', 1, 2, 4, 10),
('2021-10-11', 1, 4, 3, 15),
('2021-10-14', 2, 5, 2, 35),
('2021-10-14', 2, 2, 3, 85),
('2021-10-14', 2, 3, 4, 70),
('2021-10-14', 2, 1, 5, 63),
('2021-10-19', 1, 1, 5, 90),
('2021-10-19', 1, 2, 3, 55),
('2021-10-19', 1, 4, 5, 60),
('2021-10-21', 2, 5, 3, 34),
('2021-10-21', 2, 2, 3, 76),
('2021-10-21', 2, 3, 3, 88),
('2021-10-21', 2, 1, 5, 50),
('2021-10-22', 2, 5, 3, 40),
('2021-10-22', 2, 2, 3, 85),
('2021-10-22', 2, 3, 3, 100),
('2021-10-22', 2, 1, 5, 75);

-- --------------------------------------------------------

--
-- Table structure for table `groups`
--

CREATE TABLE `groups` (
  `id` int(11) NOT NULL,
  `name` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `groups`
--

INSERT INTO `groups` (`id`, `name`) VALUES
(1, 'Математика ЕГЭ'),
(2, 'Русский язык ЕГЭ');

-- --------------------------------------------------------

--
-- Table structure for table `last_test_res`
--

CREATE TABLE `last_test_res` (
  `idStudent` int(11) DEFAULT NULL,
  `idGroup` int(11) DEFAULT NULL,
  `res` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `last_test_res`
--

INSERT INTO `last_test_res` (`idStudent`, `idGroup`, `res`) VALUES
(1, 1, 74),
(2, 1, 68),
(4, 1, 33),
(1, 2, 45),
(2, 2, 56),
(3, 2, 62),
(5, 2, 56);

-- --------------------------------------------------------

--
-- Table structure for table `students`
--

CREATE TABLE `students` (
  `id` int(11) NOT NULL,
  `name` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `students`
--

INSERT INTO `students` (`id`, `name`) VALUES
(1, 'Алексей'),
(2, 'Мария'),
(3, 'Николай'),
(4, 'Михаил'),
(5, 'Елена');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `data_about_lessons`
--
ALTER TABLE `data_about_lessons`
  ADD KEY `idGroup` (`idGroup`),
  ADD KEY `idStudent` (`idStudent`);

--
-- Indexes for table `groups`
--
ALTER TABLE `groups`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `last_test_res`
--
ALTER TABLE `last_test_res`
  ADD KEY `idStudent` (`idStudent`) USING BTREE,
  ADD KEY `idGroup` (`idGroup`) USING BTREE;

--
-- Indexes for table `students`
--
ALTER TABLE `students`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `groups`
--
ALTER TABLE `groups`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `students`
--
ALTER TABLE `students`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `data_about_lessons`
--
ALTER TABLE `data_about_lessons`
  ADD CONSTRAINT `data_about_lessons_ibfk_1` FOREIGN KEY (`idGroup`) REFERENCES `groups` (`id`),
  ADD CONSTRAINT `data_about_lessons_ibfk_2` FOREIGN KEY (`idStudent`) REFERENCES `students` (`id`);

--
-- Constraints for table `last_test_res`
--
ALTER TABLE `last_test_res`
  ADD CONSTRAINT `last_test_res_ibfk_1` FOREIGN KEY (`idStudent`) REFERENCES `students` (`id`),
  ADD CONSTRAINT `last_test_res_ibfk_2` FOREIGN KEY (`idGroup`) REFERENCES `groups` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
