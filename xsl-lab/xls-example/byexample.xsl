<?xml version="1.0" encoding="UTF-8"?>
<!--
byexample.xsl [2021-10-02 BAR8TL]
Ejemplo de conversion de un archivo XML con datos de un empleado a un
archivo HTML
-->
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" version="1.0">
  <xsl:output method="xml" encoding="UTF-8"/>
<!-- Exemplar document -->
  <v2:employee xmlns:v1='urn:employee:v1' xmlns:v2='urn:employee:v2'>
    <name>
      <xsl:value-of select="concat(/v1:emp/fname,' ',/v1:emp/lname)"/>
    </name>
    <title>
      <xsl:value-of select='/v1:emp/position'/>
    </title>
  </v2:employee>
</xsl:stylesheet>
