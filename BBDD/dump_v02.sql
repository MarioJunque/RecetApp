-- MariaDB dump 10.17  Distrib 10.5.6-MariaDB, for osx10.15 (x86_64)
--
-- Host: localhost    Database: recetapp
-- ------------------------------------------------------
-- Server version	5.7.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `recetapp` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `recetapp`;

--
-- Table structure for table `ingrediente_receta`
--

DROP TABLE IF EXISTS `ingrediente_receta`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ingrediente_receta` (
  `id_ingrediente_receta` int(11) NOT NULL AUTO_INCREMENT,
  `id_ingredientes` int(11) DEFAULT NULL,
  `id_recetas` int(11) DEFAULT NULL,
  PRIMARY KEY (`id_ingrediente_receta`),
  KEY `id_ingrediente` (`id_ingredientes`),
  KEY `id_receta` (`id_recetas`),
  CONSTRAINT `ir1` FOREIGN KEY (`id_ingredientes`) REFERENCES `ingredientes` (`id_ingrediente`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `ir2` FOREIGN KEY (`id_recetas`) REFERENCES `recetas` (`id_receta`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ingrediente_receta`
--

LOCK TABLES `ingrediente_receta` WRITE;
/*!40000 ALTER TABLE `ingrediente_receta` DISABLE KEYS */;
/*!40000 ALTER TABLE `ingrediente_receta` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ingrediente_usuario`
--

DROP TABLE IF EXISTS `ingrediente_usuario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ingrediente_usuario` (
  `id_ingrediente_usuario` int(11) NOT NULL AUTO_INCREMENT,
  `id_ingredientes` int(11) DEFAULT NULL,
  `id_usuarios` int(11) DEFAULT NULL,
  PRIMARY KEY (`id_ingrediente_usuario`),
  KEY `id_ingrediente` (`id_ingredientes`),
  KEY `id_usuario` (`id_usuarios`),
  CONSTRAINT `c1` FOREIGN KEY (`id_ingredientes`) REFERENCES `ingredientes` (`id_ingrediente`) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `c2` FOREIGN KEY (`id_usuarios`) REFERENCES `usuarios` (`id_usuario`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ingrediente_usuario`
--

LOCK TABLES `ingrediente_usuario` WRITE;
/*!40000 ALTER TABLE `ingrediente_usuario` DISABLE KEYS */;
/*!40000 ALTER TABLE `ingrediente_usuario` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ingredientes`
--

DROP TABLE IF EXISTS `ingredientes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ingredientes` (
  `id_ingrediente` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_ingrediente`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ingredientes`
--

LOCK TABLES `ingredientes` WRITE;
/*!40000 ALTER TABLE `ingredientes` DISABLE KEYS */;
INSERT INTO `ingredientes` VALUES  (1,'guisantes'),(2,'salchichón'),(3,'atún'),(4,'orégano'),(5,'setas'),(6,'remolacha'),(7,'sandía'),(8,'salmón'),(9,'cebolleta'),(10,'langosta'),(11,'coliflor'),(12,'berenjena'),(13,'algas'),(14,'sal'),(15,'caviar'),(16,'melón'),(17,'coliflor'),(18,'jamón serrano'),(19,'calabaza'),(20,'anchoa'),(21,'chocolate'),(22,'lima'),(23,'brócoli'),(24,'gamba'),(25,'bacon'),(26,'mejillón'),(27,'plátano'),(28,'morcilla'),(29,'rodaballo'),(30,'lechuga'),(31,'nata'),(32,'ciervo'),(33,'camarón'),(34,'colorante'),(35,'melón'),(36,'perejil'),(37,'soja'),(38,'habas'),(39,'boniato'),(40,'chistorra'),(41,'mango'),(42,'cilantro'),(43,'calabacín'),(44,'lechuga'),(45,'vainilla'),(46,'frambuesa'),(47,'chile'),(48,'sésamo'),(49,'coles'),(50,'yuca'),(51,'membrillo'),(52,'ñame'),(53,'nachos'),(54,'granada'),(55,'avena'),(56,'boquerón'),(57,'pepinillo'),(58,'quinoa'),(59,'higo'),(60,'avellana'),(61,'puerro'),(62,'clavo'),(63,'lubina'),(64,'huevo'),(65,'pimiento'),
(66,'spaghetti'),(67,'fideos'),(68,'curry'),(69,'hinojo'),(70,'mafafa'),(71,'coco'),(72,'pollo'),(73,'uva'),(74,'colirrabano'),(75,'espárrago'),(76,'cardo'),(77,'mora'),(78,'endivia'),(79,'fideua'),(80,'macarrón'),(81,'whisky'),(82,'conejo'),(83,'galleta'),(84,'alcachofa'),(85,'mantequilla'),(86,'cardamomo'),(87,'rábano'),(88,'pomelo'),(89,'ketchup'),(90,'bogavante'),(91,'cúrcuma'),(92,'tomillo'),(93,'pepino'),(94,'cigala'),(95,'borraja'),(96,'cayena'),(97,'palta'),(98,'melocotón'),(99,'acelga'),(100,'nabo'),(101,'tomate'),(102,'cebolla'),(103,'zanahoria');
/*!40000 ALTER TABLE `ingredientes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `recetas`
--

DROP TABLE IF EXISTS `recetas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `recetas` (
  `id_receta` int(11) NOT NULL AUTO_INCREMENT,
  `numeroComensales` int(11) DEFAULT NULL,
  `instrucciones` varchar(800) NOT NULL,
  `nombre_receta` varchar(200) NOT NULL,
   
  PRIMARY KEY (`id_receta`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `recetas`
--

LOCK TABLES `recetas` WRITE;
/*!40000 ALTER TABLE `recetas` DISABLE KEYS */;
/*!40000 ALTER TABLE `recetas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuarios`
--

DROP TABLE IF EXISTS `usuarios`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usuarios` (
  `id_usuario` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(50) DEFAULT NULL,
  `contraseña` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
	`gluten` tinyint(1) DEFAULT 0,
	`lactosa` tinyint(1) DEFAULT 0,
	`histamina` tinyint(1) DEFAULT 0,
	PRIMARY KEY (`id_usuario`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuarios`
--

LOCK TABLES `usuarios` WRITE;
/*!40000 ALTER TABLE `usuarios` DISABLE KEYS */;
/*!40000 ALTER TABLE `usuarios` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-10-22 21:24:41
