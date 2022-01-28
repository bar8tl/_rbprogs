<?xml version="1.0" encoding="UTF-8"?>
<!--
bib.xsl [2021-10-02 BAR8TL]
Ejemplo de conversion de un archivo XML con datos de una biblioteca a un
archivo HTML
-->
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" version="1.0">
  <xsl:output method="html" version="5.0" encoding="UTF-8" indent="yes" />

<!-- Plantilla raíz -->
<xsl:template match="/">
  <html>
    <head>
    </head>
    <body>
      <xsl:apply-templates select="bib" />
    </body>
    </html>
</xsl:template>

<!-- Plantilla bib -->
<xsl:template match="/bib">
  <ul>
    <xsl:apply-templates select="book" />
  </ul>
</xsl:template>

<!-- Plantilla book -->
<xsl:template match="/bib/book">
  <li>
    <xsl:apply-templates select="title" />
    <xsl:apply-templates select="year" />
  </li>
</xsl:template>

<!-- Plantilla title -->
<xsl:template match="/bib/book/title">
  <xsl:value-of select="." />
</xsl:template>

<!-- Plantilla year -->
<xsl:template match="/bib/book/year">
  (<xsl:value-of select="." />)
</xsl:template>

</xsl:stylesheet>
