<?xml version="1.0" encoding="UTF-8"?>
<!--
#
# $Header:$
# Mapping EDIFACT DELFOR D04A -> IDOC DELFOR02
# Source description: oena_edifact_delfor97a.fde
# Target description: oena_idoc_delfor02.fde
#
# $Log:$
# date      | user                              | comment
# ..............................................................................................
# 20210806  | Guilherme Chamone SO/OPM74-BR     | Initial mapping - EDICUSTOMERREQUEST-3189
# 20210806  | Guilherme Chamone SO/OPM74-BR     | CHG33 EDK09 KTEXT_V, EDK09/EDP10 LABNK - EDICUSTOMERREQUEST-5209
# 20210910  | Guilherme Chamone SO/OPM74-BR     | Changing the map to add support for TRWAUTOELM
-->
<!--Stylesheet was automatically generated on Nov 24, 2020 10:13 AM-->
<xsl:stylesheet version="2.0"
                                xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                                xmlns:XmlTransform="java:de.axway.converter.xml.xsltextensions.XsltExtensions"
                                exclude-result-prefixes="XmlTransform"
                                xmlns:xs="http://www.w3.org/2001/XMLSchema"
                                xmlns:fn="http://www.w3.org/2005/xpath-functions"
                                xmlns:err="http://www.w3.org/2005/xqt-errors"
                                xmlns:saxon="http://saxon.sf.net/"
                                xmlns:bosch="http://www.bosch.com"
                                xmlns:exslt="http://exslt.org/common"
                                extension-element-prefixes="saxon xs fn err bosch XmlTransform exslt">

        <!--Parameters and Variables                                       -->
        <xsl:param name="ACTISEDI" select="''"/>
        <xsl:param name="PPATH" select="''"/>
        <xsl:param name="CIMTYP" select="''"/>
        <xsl:param name="IDOCTYP" select="''"/>
        <xsl:param name="MESTYP" select="''"/>
        <xsl:param name="MESCOD" select="''"/>
        <xsl:param name="MANDT" select="''"/>
        <xsl:param name="STD" select="''"/>
        <xsl:param name="STDVRS" select="''"/>
        <xsl:param name="STDMES" select="''"/>
        <xsl:param name="SNDPRN" select="''"/>
        <xsl:param name="SNDPOR" select="''"/>
        <xsl:param name="SNDPRT" select="''"/>
        <xsl:param name="SNDPFC" select="''"/>
        <xsl:param name="RCVPRN" select="''"/>
        <xsl:param name="RCVPOR" select="''"/>
        <xsl:param name="RCVPRT" select="''"/>
        <xsl:param name="RCVPFC" select="''"/>
        <xsl:param name="JOBID" select="''"/>
        <xsl:param name="PARTNER" select="''"/>
        <xsl:param name="stepfile" select="''"/>
        <xsl:variable name="gv_EARLIEST" select="''" saxon:assignable="yes"/>
        <xsl:variable name="gv_LAST" select="''" saxon:assignable="yes"/>
        <xsl:variable name="gv_DELIVERY" select="''" saxon:assignable="yes"/>
        <xsl:variable name="gv_POSEXCounter" select="0" saxon:assignable="yes"/>
        <xsl:variable name="gv_Mdate" select="''" saxon:assignable="yes"/>

        <!--SNDPRN / XmlTransform:getPropertyActisedi to use table in text format - make the conversion-->
        <xsl:variable name="gv_SNDPRN" select="XmlTransform:getPropertyActisedi('/lib/convert/tftab/OENA_CODETAB_INBOUND_PARTNER',concat(ACTIS/UNB/UNB_S002/UNB_0004,ACTIS/UNB/UNB_S003/UNB_0010,'#SNDPRN'))" saxon:assignable="yes"/>

        <!--Date, Time from XSLT 2.0 functions-->
        <xsl:variable name="gv_docdate" select="format-date(current-date(),'[Y0001][M01][D01]')"/>
        <xsl:variable name="gv_doctime" select="format-time(current-time(),'[H01][m01][s01]')"/>

        <!--START Mapping-->
        <xsl:output encoding="UTF-8" indent="yes" method="xml"/>

        <!--Create Structure-->
        <!--Template for root element-->
        <xsl:template match="/">

                <!--Write stepfile name to log file-->
                <xsl:message terminate="no">
                        <xsl:text>STEPFILE: </xsl:text>
                        <xsl:value-of select="$stepfile" />
                </xsl:message>

                <xsl:choose>
                        <xsl:when test="$gv_SNDPRN"/>
                        <xsl:otherwise>
                                <xsl:message terminate="yes">
                                        <xsl:text>ERROR: no value found in the table OENA_CODETAB_INBOUND_PARTNER for </xsl:text>
                                        <xsl:text>EDICode : </xsl:text><xsl:value-of select="concat(ACTIS/UNB/UNB_S002/UNB_0004,ACTIS/UNB/UNB_S003/UNB_0010,'#SNDPRN')" />
                                </xsl:message>
                        </xsl:otherwise>
                </xsl:choose>

                <!--The loop below allows the message to have more than one UNH-->
                <ACTIS>
                        <xsl:for-each select="ACTIS/SG0">
                                <xsl:call-template name="IDOC"/>
                        </xsl:for-each>
                </ACTIS>
        </xsl:template>

        <!--Template for 'IDOC' -->
        <xsl:template name="IDOC">
                <xsl:for-each select="SG6/SG12">
                        <IDOC>
                                <xsl:call-template name="EDI_DC" />
                                <xsl:call-template name="G01" />
                        </IDOC>
                </xsl:for-each>
        </xsl:template>

        <!--Template for 'G01' -->
        <xsl:template name="G01">
                <G01>
                        <xsl:call-template name="E2EDK09" />
                        <xsl:call-template name="G_E2EDKA1" />
                        <xsl:call-template name="G02" />
                </G01>
        </xsl:template>

        <!--Template for 'G_E2EDKA1' -->
        <xsl:template name="G_E2EDKA1">
                <G_E2EDKA1>
                        <xsl:call-template name="E2EDKA1_AG" />
                </G_E2EDKA1>
                <G_E2EDKA1>
                        <xsl:call-template name="E2EDKA1_AB" />
                </G_E2EDKA1>
                <G_E2EDKA1>
                        <xsl:call-template name="E2EDKA1_WE" />
                </G_E2EDKA1>
        </xsl:template>

        <!--Template for 'G02' -->
        <xsl:template name="G02">
                <G02>
                        <xsl:call-template name="E2EDP10" />
                        <xsl:call-template name="E2EDP15" />
                        <xsl:call-template name="E2EDP15_01" />
                        <xsl:call-template name="E2EDP15_02" />
                        <xsl:call-template name="E2EDP15_03" />
                        <xsl:call-template name="E2EDP15_04" />
                        <xsl:call-template name="E2EDP15_05" />
                        <xsl:for-each select="G_SG17/SG17_1">
                                <xsl:call-template name="E2EDP16" />
                        </xsl:for-each>
                        <xsl:for-each select="G_SG17/SG17_4">
                                <xsl:call-template name="E2EDP16" />
                        </xsl:for-each>
                </G02>
        </xsl:template>

        <!--Template for 'EDI_DC' -->
        <xsl:template name="EDI_DC">
                <EDI_DC>
                        <EDI_DC_TABNAM>
                                <xsl:value-of select="'EDI_DC40'" />
                        </EDI_DC_TABNAM>
                        <EDI_DC_MANDT>
                                <xsl:value-of select="$MANDT" />
                        </EDI_DC_MANDT>
                        <EDI_DC_DOCNUM></EDI_DC_DOCNUM>
                        <EDI_DC_DOCREL></EDI_DC_DOCREL>
                        <EDI_DC_STATUS></EDI_DC_STATUS>
                        <EDI_DC_DIRECT>
                                <xsl:value-of select="'2'" />
                        </EDI_DC_DIRECT>
                        <EDI_DC_OUTMOD></EDI_DC_OUTMOD>
                        <EDI_DC_EXPRSS></EDI_DC_EXPRSS>
                        <EDI_DC_TEST></EDI_DC_TEST>
                        <EDI_DC_IDOCTYP>
                                <xsl:value-of select="'DELFOR02'" />
                        </EDI_DC_IDOCTYP>
                        <EDI_DC_CIMTYP></EDI_DC_CIMTYP>
                        <EDI_DC_MESTYP>
                                <xsl:value-of select="'DELINS'" />
                        </EDI_DC_MESTYP>
                        <EDI_DC_MESCOD></EDI_DC_MESCOD>
                        <EDI_DC_MESFCT></EDI_DC_MESFCT>
                        <EDI_DC_STD>
                                <xsl:value-of select="'E'" />
                        </EDI_DC_STD>
                        <EDI_DC_STDVRS>
                                <xsl:value-of select="concat(../../UNH/UNH_S009/UNH_0052,../../UNH/UNH_S009/UNH_0054)" />
                        </EDI_DC_STDVRS>
                        <EDI_DC_STDMES>
                                <xsl:value-of select="../../UNH/UNH_S009/UNH_0065" />
                        </EDI_DC_STDMES>
                        <EDI_DC_SNDPOR>
                                <xsl:value-of select="$SNDPOR" />
                        </EDI_DC_SNDPOR>
                        <EDI_DC_SNDPRT>
                                <xsl:value-of select="$SNDPRT" />
                        </EDI_DC_SNDPRT>
                        <EDI_DC_SNDPFC>
                                <xsl:value-of select="$SNDPFC" />
                        </EDI_DC_SNDPFC>
                        <EDI_DC_SNDPRN>
                                <xsl:value-of select="$gv_SNDPRN" />
                        </EDI_DC_SNDPRN>
                        <EDI_DC_SNDSAD></EDI_DC_SNDSAD>
                        <EDI_DC_SNDLAD>
                                <xsl:value-of select="concat(../../../UNB/UNB_S002/UNB_0007_01,'*',../../../UNB/UNB_S002/UNB_0004)" />
                        </EDI_DC_SNDLAD>
                        <EDI_DC_RCVPOR>
                                <xsl:value-of select="$RCVPOR" />
                        </EDI_DC_RCVPOR>
                        <EDI_DC_RCVPRT>
                                <xsl:value-of select="$RCVPRT" />
                        </EDI_DC_RCVPRT>
                        <EDI_DC_RCVPFC>
                                <xsl:value-of select="$RCVPFC" />
                        </EDI_DC_RCVPFC>
                        <EDI_DC_RCVPRN>
                                <xsl:value-of select="$RCVPRN" />
                        </EDI_DC_RCVPRN>
                        <EDI_DC_RCVSAD></EDI_DC_RCVSAD>
                        <EDI_DC_RCVLAD>
                                <xsl:value-of select="concat(../../../UNB/UNB_S003/UNB_0007_02,'*',../../../UNB/UNB_S003/UNB_0010)" />
                        </EDI_DC_RCVLAD>
                        <EDI_DC_CREDAT>
                                <xsl:value-of select="$gv_docdate" />
                        </EDI_DC_CREDAT>
                        <EDI_DC_CRETIM>
                                <xsl:value-of select="$gv_doctime" />
                        </EDI_DC_CRETIM>
                        <EDI_DC_REFINT>
                                <xsl:value-of select="../../../UNB/UNB_0020"/>
                        </EDI_DC_REFINT>
                        <EDI_DC_REFGRP>
                                <xsl:value-of select="../../UNH/UNH_0062"/>
                        </EDI_DC_REFGRP>
                        <EDI_DC_REFMES></EDI_DC_REFMES>
                        <EDI_DC_ARCKEY>
                                <xsl:value-of select="concat('workflow-ID: ',$JOBID)"/>
                        </EDI_DC_ARCKEY>
                        <EDI_DC_SERIAL></EDI_DC_SERIAL>
                </EDI_DC>
        </xsl:template>

        <!--Template for 'E2EDK09' -->
        <xsl:template name="E2EDK09">
                <E2EDK09>
                        <E2EDK09_SEGNAM></E2EDK09_SEGNAM>
                        <E2EDK09_MANDT></E2EDK09_MANDT>
                        <E2EDK09_DOCNUM></E2EDK09_DOCNUM>
                        <E2EDK09_SEGNUM></E2EDK09_SEGNUM>
                        <E2EDK09_PSGNUM></E2EDK09_PSGNUM>
                        <E2EDK09_HLEVEL></E2EDK09_HLEVEL>
                        <E2EDK09_VTRNR>
                                <xsl:value-of select="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'ON']/RFF_04_1154"/>
                        </E2EDK09_VTRNR>
                        <E2EDK09_BSTDK>
                                <xsl:choose>
                                        <xsl:when test="../../DTM_01_gr/DTM_01_137/DTM_01_137_C507[DTM_01_137_2005 = '137']/DTM_01_137_2380 != ''">
                                                <xsl:value-of select="../../DTM_01_gr/DTM_01_137/DTM_01_137_C507[DTM_01_137_2005 = '137']/DTM_01_137_2380" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="$gv_docdate" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDK09_BSTDK>
                        <E2EDK09_LABNK>
                                <xsl:choose>
                                        <xsl:when test="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'AAN']/RFF_04_1154 != ''">
                                                <xsl:value-of select="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'AAN']/RFF_04_1154"/>
                                        </xsl:when>
                                        <xsl:when test="../../BGM[BGM_C002/BGM_1001 = '241']/BGM_C106/BGM_1004 != ''">
                                                <xsl:value-of select="../../BGM[BGM_C002/BGM_1001 = '241']/BGM_C106/BGM_1004" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="concat($gv_docdate, substring($gv_doctime, 1, 4))" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDK09_LABNK>
                        <E2EDK09_ZEICH></E2EDK09_ZEICH>
                        <E2EDK09_BSTZD></E2EDK09_BSTZD>
                        <E2EDK09_ABRVW>
                                <xsl:value-of select="'S'" />
                        </E2EDK09_ABRVW>
                        <E2EDK09_KSTAT></E2EDK09_KSTAT>
                        <E2EDK09_KTEXT></E2EDK09_KTEXT>
                        <E2EDK09_ABNRA>
                                <xsl:value-of select="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'AIF']/RFF_04_1154" />
                        </E2EDK09_ABNRA>
                        <E2EDK09_ABNRD>
                                <xsl:for-each select="SG13">
                                        <xsl:if test="RFF_04/RFF_04_C506/RFF_04_1153 = 'AIF'">
                                                <xsl:value-of select="DTM_09/DTM_09_C507[DTM_09_2005 = '137']/DTM_09_2380" />
                                        </xsl:if>
                                </xsl:for-each>
                        </E2EDK09_ABNRD>
                        <E2EDK09_KTEXT_V>
                                <xsl:choose>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'SE']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SE']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDK09_KTEXT_V>
                        <E2EDK09_USR01>
                                <xsl:for-each select="G_SG17/SG17_2">
                                        <xsl:choose>
                                                <xsl:when test="SG18_2/QTY_02/QTY_02_C186/QTY_02_6060 != ''">
                                                        <xsl:value-of select="SG18_2/QTY_02/QTY_02_C186/QTY_02_6060" />
                                                </xsl:when>
                                        </xsl:choose>
                                </xsl:for-each>
                        </E2EDK09_USR01>
                        <E2EDK09_USR02></E2EDK09_USR02>
                        <E2EDK09_USR03>
                                <xsl:for-each select="G_SG17/SG17_3">
                                        <xsl:choose>
                                                <xsl:when test="SG18_3/QTY_02/QTY_02_C186/QTY_02_6060 != ''">
                                                        <xsl:value-of select="SG18_3/QTY_02/QTY_02_C186/QTY_02_6060" />
                                                </xsl:when>
                                        </xsl:choose>
                                </xsl:for-each>
                        </E2EDK09_USR03>
                        <E2EDK09_USR04></E2EDK09_USR04>
                        <E2EDK09_USR05></E2EDK09_USR05>
                        <E2EDK09_CYEFZ></E2EDK09_CYEFZ>
                        <E2EDK09_CYDAT></E2EDK09_CYDAT>
                        <E2EDK09_MFLAUF></E2EDK09_MFLAUF>
                        <E2EDK09_MFEIN></E2EDK09_MFEIN>
                        <E2EDK09_FFLAUF></E2EDK09_FFLAUF>
                        <E2EDK09_FFEIN></E2EDK09_FFEIN>
                        <rest></rest>
                </E2EDK09>
        </xsl:template>

        <!--Template for 'E2EDKA1_AG' -->
        <xsl:template name="E2EDKA1_AG">
                <E2EDKA1_AG>
                        <E2EDKA1_AG_SEGNAM></E2EDKA1_AG_SEGNAM>
                        <E2EDKA1_AG_MANDT></E2EDKA1_AG_MANDT>
                        <E2EDKA1_AG_DOCNUM></E2EDKA1_AG_DOCNUM>
                        <E2EDKA1_AG_SEGNUM></E2EDKA1_AG_SEGNUM>
                        <E2EDKA1_AG_PSGNUM></E2EDKA1_AG_PSGNUM>
                        <E2EDKA1_AG_HLEVEL></E2EDKA1_AG_HLEVEL>
                        <E2EDKA1_AG_PARVW>
                                <xsl:value-of select="'AG'" />
                        </E2EDKA1_AG_PARVW>
                        <E2EDKA1_AG_PARTN>
                                <xsl:choose>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'MI']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'MI']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'BY']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'BY']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDKA1_AG_PARTN>
                        <E2EDKA1_AG_LIFNR>
                                <xsl:choose>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'SE']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SE']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDKA1_AG_LIFNR>
                        <E2EDKA1_AG_NAME1></E2EDKA1_AG_NAME1>
                        <E2EDKA1_AG_NAME2></E2EDKA1_AG_NAME2>
                        <E2EDKA1_AG_NAME3></E2EDKA1_AG_NAME3>
                        <E2EDKA1_AG_NAME4></E2EDKA1_AG_NAME4>
                        <E2EDKA1_AG_STRAS></E2EDKA1_AG_STRAS>
                        <E2EDKA1_AG_STRS2></E2EDKA1_AG_STRS2>
                        <E2EDKA1_AG_PFACH></E2EDKA1_AG_PFACH>
                        <E2EDKA1_AG_ORT01></E2EDKA1_AG_ORT01>
                        <E2EDKA1_AG_COUNC></E2EDKA1_AG_COUNC>
                        <E2EDKA1_AG_PSTLZ></E2EDKA1_AG_PSTLZ>
                        <E2EDKA1_AG_PSTL2></E2EDKA1_AG_PSTL2>
                        <E2EDKA1_AG_LAND1></E2EDKA1_AG_LAND1>
                        <E2EDKA1_AG_ABLAD></E2EDKA1_AG_ABLAD>
                        <E2EDKA1_AG_PERNR></E2EDKA1_AG_PERNR>
                        <E2EDKA1_AG_PARNR></E2EDKA1_AG_PARNR>
                        <E2EDKA1_AG_TELF1></E2EDKA1_AG_TELF1>
                        <E2EDKA1_AG_TELF2></E2EDKA1_AG_TELF2>
                        <E2EDKA1_AG_TELBX></E2EDKA1_AG_TELBX>
                        <E2EDKA1_AG_TELFX></E2EDKA1_AG_TELFX>
                        <E2EDKA1_AG_TELTX></E2EDKA1_AG_TELTX>
                        <E2EDKA1_AG_TELX1></E2EDKA1_AG_TELX1>
                        <E2EDKA1_AG_SPRAS></E2EDKA1_AG_SPRAS>
                        <E2EDKA1_AG_ANRED></E2EDKA1_AG_ANRED>
                        <E2EDKA1_AG_ORT02></E2EDKA1_AG_ORT02>
                        <E2EDKA1_AG_HAUSN></E2EDKA1_AG_HAUSN>
                        <E2EDKA1_AG_STOCK></E2EDKA1_AG_STOCK>
                        <E2EDKA1_AG_REGIO></E2EDKA1_AG_REGIO>
                        <E2EDKA1_AG_PARGE></E2EDKA1_AG_PARGE>
                        <E2EDKA1_AG_ISOAL></E2EDKA1_AG_ISOAL>
                        <E2EDKA1_AG_ISONU></E2EDKA1_AG_ISONU>
                        <E2EDKA1_AG_FCODE></E2EDKA1_AG_FCODE>
                        <E2EDKA1_AG_IHREZ></E2EDKA1_AG_IHREZ>
                        <E2EDKA1_AG_BNAME></E2EDKA1_AG_BNAME>
                        <E2EDKA1_AG_PAORG></E2EDKA1_AG_PAORG>
                        <E2EDKA1_AG_ORGTX></E2EDKA1_AG_ORGTX>
                        <E2EDKA1_AG_PAGRU></E2EDKA1_AG_PAGRU>
                        <E2EDKA1_AG_KNREF></E2EDKA1_AG_KNREF>
                        <E2EDKA1_AG_ILNNR></E2EDKA1_AG_ILNNR>
                        <E2EDKA1_AG_PFORT></E2EDKA1_AG_PFORT>
                        <E2EDKA1_AG_SPRAS_ISO></E2EDKA1_AG_SPRAS_ISO>
                        <REST></REST>
                </E2EDKA1_AG>
        </xsl:template>

        <!--Template for 'E2EDKA1_AB' -->
        <xsl:template name="E2EDKA1_AB">
                <E2EDKA1_AB>
                        <E2EDKA1_AB_SEGNAM></E2EDKA1_AB_SEGNAM>
                        <E2EDKA1_AB_MANDT></E2EDKA1_AB_MANDT>
                        <E2EDKA1_AB_DOCNUM></E2EDKA1_AB_DOCNUM>
                        <E2EDKA1_AB_SEGNUM></E2EDKA1_AB_SEGNUM>
                        <E2EDKA1_AB_PSGNUM></E2EDKA1_AB_PSGNUM>
                        <E2EDKA1_AB_HLEVEL></E2EDKA1_AB_HLEVEL>
                        <E2EDKA1_AB_PARVW>
                                <xsl:value-of select="'AB'" />
                        </E2EDKA1_AB_PARVW>
                        <E2EDKA1_AB_PARTN>
                                <xsl:value-of select="LOC_02[LOC_02_3227 = '11']/LOC_02_C517/LOC_02_3225" />
                        </E2EDKA1_AB_PARTN>
                        <E2EDKA1_AB_LIFNR></E2EDKA1_AB_LIFNR>
                        <E2EDKA1_AB_NAME1></E2EDKA1_AB_NAME1>
                        <E2EDKA1_AB_NAME2></E2EDKA1_AB_NAME2>
                        <E2EDKA1_AB_NAME3></E2EDKA1_AB_NAME3>
                        <E2EDKA1_AB_NAME4></E2EDKA1_AB_NAME4>
                        <E2EDKA1_AB_STRAS></E2EDKA1_AB_STRAS>
                        <E2EDKA1_AB_STRS2></E2EDKA1_AB_STRS2>
                        <E2EDKA1_AB_PFACH></E2EDKA1_AB_PFACH>
                        <E2EDKA1_AB_ORT01></E2EDKA1_AB_ORT01>
                        <E2EDKA1_AB_COUNC></E2EDKA1_AB_COUNC>
                        <E2EDKA1_AB_PSTLZ></E2EDKA1_AB_PSTLZ>
                        <E2EDKA1_AB_PSTL2></E2EDKA1_AB_PSTL2>
                        <E2EDKA1_AB_LAND1></E2EDKA1_AB_LAND1>
                        <E2EDKA1_AB_ABLAD></E2EDKA1_AB_ABLAD>
                        <E2EDKA1_AB_PERNR></E2EDKA1_AB_PERNR>
                        <E2EDKA1_AB_PARNR></E2EDKA1_AB_PARNR>
                        <E2EDKA1_AB_TELF1></E2EDKA1_AB_TELF1>
                        <E2EDKA1_AB_TELF2></E2EDKA1_AB_TELF2>
                        <E2EDKA1_AB_TELBX></E2EDKA1_AB_TELBX>
                        <E2EDKA1_AB_TELFX></E2EDKA1_AB_TELFX>
                        <E2EDKA1_AB_TELTX></E2EDKA1_AB_TELTX>
                        <E2EDKA1_AB_TELX1></E2EDKA1_AB_TELX1>
                        <E2EDKA1_AB_SPRAS></E2EDKA1_AB_SPRAS>
                        <E2EDKA1_AB_ANRED></E2EDKA1_AB_ANRED>
                        <E2EDKA1_AB_ORT02></E2EDKA1_AB_ORT02>
                        <E2EDKA1_AB_HAUSN></E2EDKA1_AB_HAUSN>
                        <E2EDKA1_AB_STOCK></E2EDKA1_AB_STOCK>
                        <E2EDKA1_AB_REGIO></E2EDKA1_AB_REGIO>
                        <E2EDKA1_AB_PARGE></E2EDKA1_AB_PARGE>
                        <E2EDKA1_AB_ISOAL></E2EDKA1_AB_ISOAL>
                        <E2EDKA1_AB_ISONU></E2EDKA1_AB_ISONU>
                        <E2EDKA1_AB_FCODE></E2EDKA1_AB_FCODE>
                        <E2EDKA1_AB_IHREZ></E2EDKA1_AB_IHREZ>
                        <E2EDKA1_AB_BNAME></E2EDKA1_AB_BNAME>
                        <E2EDKA1_AB_PAORG></E2EDKA1_AB_PAORG>
                        <E2EDKA1_AB_ORGTX></E2EDKA1_AB_ORGTX>
                        <E2EDKA1_AB_PAGRU></E2EDKA1_AB_PAGRU>
                        <E2EDKA1_AB_KNREF></E2EDKA1_AB_KNREF>
                        <E2EDKA1_AB_ILNNR></E2EDKA1_AB_ILNNR>
                        <E2EDKA1_AB_PFORT></E2EDKA1_AB_PFORT>
                        <E2EDKA1_AB_SPRAS_ISO></E2EDKA1_AB_SPRAS_ISO>
                        <REST></REST>
                </E2EDKA1_AB>
        </xsl:template>

        <!--Template for 'E2EDKA1_WE' -->
        <xsl:template name="E2EDKA1_WE">
                <E2EDKA1_WE>
                        <E2EDKA1_WE_SEGNAM></E2EDKA1_WE_SEGNAM>
                        <E2EDKA1_WE_MANDT></E2EDKA1_WE_MANDT>
                        <E2EDKA1_WE_DOCNUM></E2EDKA1_WE_DOCNUM>
                        <E2EDKA1_WE_SEGNUM></E2EDKA1_WE_SEGNUM>
                        <E2EDKA1_WE_PSGNUM></E2EDKA1_WE_PSGNUM>
                        <E2EDKA1_WE_HLEVEL></E2EDKA1_WE_HLEVEL>
                        <E2EDKA1_WE_PARVW>
                                <xsl:value-of select="'WE'" />
                        </E2EDKA1_WE_PARVW>
                        <E2EDKA1_WE_PARTN>
                                <xsl:choose>
                                        <xsl:when test="../SG7/NAD_02[NAD_02_3035 = 'ST']/NAD_02_C082/NAD_02_3039 !=''">
                                                <xsl:value-of select="../SG7/NAD_02[NAD_02_3035 = 'ST']/NAD_02_C082/NAD_02_3039" />
                                        </xsl:when>
                                        <xsl:when test="../SG7/NAD_02[NAD_02_3035 = 'DP']/NAD_02_C082/NAD_02_3039 !=''">
                                                <xsl:value-of select="../SG7/NAD_02[NAD_02_3035 = 'DP']/NAD_02_C082/NAD_02_3039" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="../../../UNB/UNB_S002/UNB_0004" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDKA1_WE_PARTN>
                        <E2EDKA1_WE_LIFNR></E2EDKA1_WE_LIFNR>
                        <E2EDKA1_WE_NAME1></E2EDKA1_WE_NAME1>
                        <E2EDKA1_WE_NAME2></E2EDKA1_WE_NAME2>
                        <E2EDKA1_WE_NAME3></E2EDKA1_WE_NAME3>
                        <E2EDKA1_WE_NAME4></E2EDKA1_WE_NAME4>
                        <E2EDKA1_WE_STRAS></E2EDKA1_WE_STRAS>
                        <E2EDKA1_WE_STRS2></E2EDKA1_WE_STRS2>
                        <E2EDKA1_WE_PFACH></E2EDKA1_WE_PFACH>
                        <E2EDKA1_WE_ORT01></E2EDKA1_WE_ORT01>
                        <E2EDKA1_WE_COUNC></E2EDKA1_WE_COUNC>
                        <E2EDKA1_WE_PSTLZ></E2EDKA1_WE_PSTLZ>
                        <E2EDKA1_WE_PSTL2></E2EDKA1_WE_PSTL2>
                        <E2EDKA1_WE_LAND1></E2EDKA1_WE_LAND1>
                        <E2EDKA1_WE_OBLAD></E2EDKA1_WE_OBLAD>
                        <E2EDKA1_WE_PERNR></E2EDKA1_WE_PERNR>
                        <E2EDKA1_WE_PARNR></E2EDKA1_WE_PARNR>
                        <E2EDKA1_WE_TELF1></E2EDKA1_WE_TELF1>
                        <E2EDKA1_WE_TELF2></E2EDKA1_WE_TELF2>
                        <E2EDKA1_WE_TELBX></E2EDKA1_WE_TELBX>
                        <E2EDKA1_WE_TELFX></E2EDKA1_WE_TELFX>
                        <E2EDKA1_WE_TELTX></E2EDKA1_WE_TELTX>
                        <E2EDKA1_WE_TELX1></E2EDKA1_WE_TELX1>
                        <E2EDKA1_WE_SPRAS></E2EDKA1_WE_SPRAS>
                        <E2EDKA1_WE_ANRED></E2EDKA1_WE_ANRED>
                        <E2EDKA1_WE_ORT02></E2EDKA1_WE_ORT02>
                        <E2EDKA1_WE_HAUSN></E2EDKA1_WE_HAUSN>
                        <E2EDKA1_WE_STOCK></E2EDKA1_WE_STOCK>
                        <E2EDKA1_WE_REGIO></E2EDKA1_WE_REGIO>
                        <E2EDKA1_WE_PARGE></E2EDKA1_WE_PARGE>
                        <E2EDKA1_WE_ISOAL></E2EDKA1_WE_ISOAL>
                        <E2EDKA1_WE_ISONU></E2EDKA1_WE_ISONU>
                        <E2EDKA1_WE_FCODE></E2EDKA1_WE_FCODE>
                        <E2EDKA1_WE_IHREZ></E2EDKA1_WE_IHREZ>
                        <E2EDKA1_WE_BNAME></E2EDKA1_WE_BNAME>
                        <E2EDKA1_WE_PAORG></E2EDKA1_WE_PAORG>
                        <E2EDKA1_WE_ORGTX></E2EDKA1_WE_ORGTX>
                        <E2EDKA1_WE_PAGRU></E2EDKA1_WE_PAGRU>
                        <E2EDKA1_WE_KNREF></E2EDKA1_WE_KNREF>
                        <E2EDKA1_WE_ILNNR></E2EDKA1_WE_ILNNR>
                        <E2EDKA1_WE_PFORT></E2EDKA1_WE_PFORT>
                        <E2EDKA1_WE_SPRAS_ISO></E2EDKA1_WE_SPRAS_ISO>
                        <REST></REST>
                </E2EDKA1_WE>
        </xsl:template>

        <!--Template for 'E2EDKA1' -->
        <xsl:template name="E2EDKA1">
                <E2EDKA1>
                        <E2EDKA1_SEGNAM></E2EDKA1_SEGNAM>
                        <E2EDKA1_MANDT></E2EDKA1_MANDT>
                        <E2EDKA1_DOCNUM></E2EDKA1_DOCNUM>
                        <E2EDKA1_SEGNUM></E2EDKA1_SEGNUM>
                        <E2EDKA1_PSGNUM></E2EDKA1_PSGNUM>
                        <E2EDKA1_HLEVEL></E2EDKA1_HLEVEL>
                        <E2EDKA1_PARVW></E2EDKA1_PARVW>
                        <E2EDKA1_PARTN></E2EDKA1_PARTN>
                        <E2EDKA1_LIFNR></E2EDKA1_LIFNR>
                        <E2EDKA1_NAME1></E2EDKA1_NAME1>
                        <E2EDKA1_NAME2></E2EDKA1_NAME2>
                        <E2EDKA1_NAME3></E2EDKA1_NAME3>
                        <E2EDKA1_NAME4></E2EDKA1_NAME4>
                        <E2EDKA1_STRAS></E2EDKA1_STRAS>
                        <E2EDKA1_STRS2></E2EDKA1_STRS2>
                        <E2EDKA1_PFACH></E2EDKA1_PFACH>
                        <E2EDKA1_ORT01></E2EDKA1_ORT01>
                        <E2EDKA1_COUNC></E2EDKA1_COUNC>
                        <E2EDKA1_PSTLZ></E2EDKA1_PSTLZ>
                        <E2EDKA1_PSTL2></E2EDKA1_PSTL2>
                        <E2EDKA1_LAND1></E2EDKA1_LAND1>
                        <E2EDKA1_ABLAD></E2EDKA1_ABLAD>
                        <E2EDKA1_PERNR></E2EDKA1_PERNR>
                        <E2EDKA1_PARNR></E2EDKA1_PARNR>
                        <E2EDKA1_TELF1></E2EDKA1_TELF1>
                        <E2EDKA1_TELF2></E2EDKA1_TELF2>
                        <E2EDKA1_TELBX></E2EDKA1_TELBX>
                        <E2EDKA1_TELFX></E2EDKA1_TELFX>
                        <E2EDKA1_TELTX></E2EDKA1_TELTX>
                        <E2EDKA1_TELX1></E2EDKA1_TELX1>
                        <E2EDKA1_SPRAS></E2EDKA1_SPRAS>
                        <E2EDKA1_ANRED></E2EDKA1_ANRED>
                        <E2EDKA1_ORT02></E2EDKA1_ORT02>
                        <E2EDKA1_HAUSN></E2EDKA1_HAUSN>
                        <E2EDKA1_STOCK></E2EDKA1_STOCK>
                        <E2EDKA1_REGIO></E2EDKA1_REGIO>
                        <E2EDKA1_PARGE></E2EDKA1_PARGE>
                        <E2EDKA1_ISOAL></E2EDKA1_ISOAL>
                        <E2EDKA1_ISONU></E2EDKA1_ISONU>
                        <E2EDKA1_FCODE></E2EDKA1_FCODE>
                        <E2EDKA1_IHREZ></E2EDKA1_IHREZ>
                        <E2EDKA1_BNAME></E2EDKA1_BNAME>
                        <E2EDKA1_PAORG></E2EDKA1_PAORG>
                        <E2EDKA1_ORGTX></E2EDKA1_ORGTX>
                        <E2EDKA1_PAGRU></E2EDKA1_PAGRU>
                        <E2EDKA1_KNREF></E2EDKA1_KNREF>
                        <E2EDKA1_ILNNR></E2EDKA1_ILNNR>
                        <E2EDKA1_PFORT></E2EDKA1_PFORT>
                        <E2EDKA1_SPRAS_ISO></E2EDKA1_SPRAS_ISO>
                        <REST></REST>
                </E2EDKA1>
        </xsl:template>

        <!--Template for 'E2EDP10' -->
        <xsl:template name="E2EDP10">
                <E2EDP10>
                        <E2EDP10_SEGNAM></E2EDP10_SEGNAM>
                        <E2EDP10_MANDT></E2EDP10_MANDT>
                        <E2EDP10_DOCNUM></E2EDP10_DOCNUM>
                        <E2EDP10_SEGNUM></E2EDP10_SEGNUM>
                        <E2EDP10_PSGNUM></E2EDP10_PSGNUM>
                        <E2EDP10_HLEVEL></E2EDP10_HLEVEL>
                        <E2EDP10_IDNKD>
                                <xsl:choose>
                                        <xsl:when test="LIN/LIN_C212[LIN_7143 = 'IN']/LIN_7140 != ''">
                                                <xsl:value-of select="LIN/LIN_C212[LIN_7143 = 'IN']/LIN_7140" />
                                        </xsl:when>
                                        <xsl:when test="LIN/LIN_C212[LIN_7143 = 'BP']/LIN_7140 != ''">
                                                <xsl:value-of select="LIN/LIN_C212[LIN_7143 = 'BP']/LIN_7140" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP10_IDNKD>
                        <E2EDP10_ARKTX>
                                <xsl:value-of select="IMD/IMD_C273/IMD_7008_01" />
                        </E2EDP10_ARKTX>
                        <E2EDP10_VRKME></E2EDP10_VRKME>
                        <E2EDP10_KWERK>
                                <xsl:choose>
                                        <xsl:when test="../SG7/NAD_02[NAD_02_3035 = 'ST']/NAD_02_C082/NAD_02_3039 !=''">
                                                <xsl:value-of select="../SG7/NAD_02[NAD_02_3035 = 'ST']/NAD_02_C082/NAD_02_3039" />
                                        </xsl:when>
                                        <xsl:when test="../SG7/NAD_02[NAD_02_3035 = 'DP']/NAD_02_C082/NAD_02_3039 !=''">
                                                <xsl:value-of select="../SG7/NAD_02[NAD_02_3035 = 'DP']/NAD_02_C082/NAD_02_3039" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="../../../UNB/UNB_S002/UNB_0004" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDP10_KWERK>
                        <E2EDP10_KLGOR></E2EDP10_KLGOR>
                        <E2EDP10_DFABL>
                                <xsl:value-of select="LOC_02[LOC_02_3227 = '11']/LOC_02_C517/LOC_02_3225" />
                        </E2EDP10_DFABL>
                        <E2EDP10_VBRST>
                                <xsl:value-of select="LOC_02[LOC_02_3227 = '159']/LOC_02_C517/LOC_02_3225" />
                        </E2EDP10_VBRST>
                        <E2EDP10_BELNR>
                                <xsl:choose>
                                        <xsl:when test="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'SI']/RFF_04_1154 != ''">
                                                <xsl:value-of select="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'SI']/RFF_04_1154"/>
                                        </xsl:when>
                                        <xsl:when test="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'AAK']/RFF_04_1154 != ''">
                                                <xsl:value-of select="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'AAK']/RFF_04_1154"/>
                                        </xsl:when>
                                        <xsl:when test="SG15/SG16/RFF_05/RFF_05_C506[RFF_05_1153 = 'SI']/RFF_05_1154 != ''">
                                                <xsl:value-of select="SG15/SG16/RFF_05/RFF_05_C506[RFF_05_1153 = 'SI']/RFF_05_1154"/>
                                        </xsl:when>
                                        <xsl:when test="SG15/SG16/RFF_05/RFF_05_C506[RFF_05_1153 = 'AAK']/RFF_05_1154 != ''">
                                                <xsl:value-of select="SG15/SG16/RFF_05/RFF_05_C506[RFF_05_1153 = 'AAK']/RFF_05_1154"/>
                                        </xsl:when>
                                </xsl:choose> =
                        </E2EDP10_BELNR>
                        <E2EDP10_LFIMG></E2EDP10_LFIMG>
                        <E2EDP10_VEMNG>
                                <xsl:value-of select="SG15/QTY_01/QTY_01_C186[QTY_01_6063 = '48']/QTY_01_6060" />
                        </E2EDP10_VEMNG>
                        <E2EDP10_LIDTL>
                                <xsl:value-of select="SG15/DTM_11/DTM_11_C507[DTM_11_2005 = '50']/DTM_11_2380" />
                        </E2EDP10_LIDTL>
                        <E2EDP10_LIFST></E2EDP10_LIFST>
                        <E2EDP10_KRITB></E2EDP10_KRITB>
                        <E2EDP10_ABHOR></E2EDP10_ABHOR>
                        <E2EDP10_ABDAT></E2EDP10_ABDAT>
                        <E2EDP10_FZDIF></E2EDP10_FZDIF>
                        <E2EDP10_AKUBM></E2EDP10_AKUBM>
                        <E2EDP10_AKUEM>
                                <xsl:value-of select="SG15/QTY_01/QTY_01_C186[QTY_01_6063 = '70']/QTY_01_6060" />
                        </E2EDP10_AKUEM>
                        <E2EDP10_DESRE></E2EDP10_DESRE>
                        <E2EDP10_ECHNO></E2EDP10_ECHNO>
                        <E2EDP10_ECHDT></E2EDP10_ECHDT>
                        <E2EDP10_ABFDE></E2EDP10_ABFDE>
                        <E2EDP10_MFADT></E2EDP10_MFADT>
                        <E2EDP10_ABMDE></E2EDP10_ABMDE>
                        <E2EDP10_FPSDN></E2EDP10_FPSDN>
                        <E2EDP10_LABKY>
                                <xsl:value-of select="'2'" />
                        </E2EDP10_LABKY>
                        <E2EDP10_TSTKY></E2EDP10_TSTKY>
                        <E2EDP10_DOKPF></E2EDP10_DOKPF>
                        <E2EDP10_KZAUS></E2EDP10_KZAUS>
                        <E2EDP10_ABRAB>
                                <xsl:choose>
                                        <xsl:when test="SG15[QTY_01/QTY_01_C186/QTY_01_6063 = '70']/DTM_11/DTM_11_C507[DTM_11_2005 = '51']/DTM_11_2380 != ''">
                                                <xsl:value-of select="SG15[QTY_01/QTY_01_C186/QTY_01_6063 = '70']/DTM_11/DTM_11_C507[DTM_11_2005 = '51']/DTM_11_2380" />
                                        </xsl:when>
                                        <xsl:when test="SG15[QTY_01/QTY_01_C186/QTY_01_6063 = '70']/DTM_11/DTM_11_C507[DTM_11_2005 = '51']/DTM_11_2380 = ''">
                                                <xsl:value-of select="'00000000'" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="''" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDP10_ABRAB>
                        <E2EDP10_ABRBI>
                                <xsl:choose>
                                        <xsl:when test="SG15[QTY_01/QTY_01_C186/QTY_01_6063 = '70']/DTM_11/DTM_11_C507[DTM_11_2005 = '52']/DTM_11_2380 != ''">
                                                <xsl:value-of select="SG15[QTY_01/QTY_01_C186/QTY_01_6063 = '70']/DTM_11/DTM_11_C507[DTM_11_2005 = '51']/DTM_11_2380" />
                                        </xsl:when>
                                        <xsl:when test="SG15[QTY_01/QTY_01_C186/QTY_01_6063 = '70']/DTM_11/DTM_11_C507[DTM_11_2005 = '52']/DTM_11_2380 = ''">
                                                <xsl:value-of select="$gv_docdate" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="''" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDP10_ABRBI>
                        <E2EDP10_SCREL>
                                <xsl:value-of select="'03'" />
                        </E2EDP10_SCREL>
                        <E2EDP10_NULDT></E2EDP10_NULDT>
                        <E2EDP10_TEART></E2EDP10_TEART>
                        <E2EDP10_SOLLFZ></E2EDP10_SOLLFZ>
                        <E2EDP10_SOLLDT></E2EDP10_SOLLDT>
                        <E2EDP10_IDNLF></E2EDP10_IDNLF>
                        <E2EDP10_BSTDK></E2EDP10_BSTDK>
                        <E2EDP10_LABNK>
                                <xsl:choose>
                                        <xsl:when test="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'AAN']/RFF_04_1154 != ''">
                                                <xsl:value-of select="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'AAN']/RFF_04_1154"/>
                                        </xsl:when>
                                        <xsl:when test="../../BGM[BGM_C002/BGM_1001 = '241']/BGM_C106/BGM_1004 != ''">
                                                <xsl:value-of select="../../BGM[BGM_C002/BGM_1001 = '241']/BGM_C106/BGM_1004" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="concat($gv_docdate, substring($gv_doctime, 1, 4))" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDP10_LABNK>
                        <E2EDP10_ABNRA></E2EDP10_ABNRA>
                        <E2EDP10_ABNRD>
                                <xsl:value-of select="SG13[RFF_04/RFF_04_C506/RFF_04_1153 = 'AAN']/DTM_09/DTM_09_C507[DTM_09_2005 = '137']/DTM_09_2380"/>
                        </E2EDP10_ABNRD>
                        <E2EDP10_VTRNR></E2EDP10_VTRNR>
                        <E2EDP10_MFRFZ></E2EDP10_MFRFZ>
                        <E2EDP10_POSEX>
                                <xsl:choose>
                                        <xsl:when test="LIN/LIN_1082_01 != ''">
                                                <xsl:value-of select="LIN/LIN_1082_01"/>
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="'1'" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDP10_POSEX>
                        <E2EDP10_IDNKD_EXTERNAL></E2EDP10_IDNKD_EXTERNAL>
                        <E2EDP10_IDNKD_GUID></E2EDP10_IDNKD_GUID>
                        <E2EDP10_IDNKD_VERSION></E2EDP10_IDNKD_VERSION>
                        <rest></rest>
                </E2EDP10>
        </xsl:template>

        <!--Template for 'E2EDP15' -->
        <xsl:template name="E2EDP15">
                <E2EDP15>
                        <E2EDP15_SEGNAM></E2EDP15_SEGNAM>
                        <E2EDP15_MANDT></E2EDP15_MANDT>
                        <E2EDP15_DOCNUM></E2EDP15_DOCNUM>
                        <E2EDP15_SEGNUM></E2EDP15_SEGNUM>
                        <E2EDP15_PSGNUM></E2EDP15_PSGNUM>
                        <E2EDP15_HLEVEL></E2EDP15_HLEVEL>
                        <E2EDP15_TDNAME>
                                <xsl:value-of select="'YAQU'" />
                        </E2EDP15_TDNAME>
                        <E2EDP15_SPRAS></E2EDP15_SPRAS>
                        <E2EDP15_TXTLF></E2EDP15_TXTLF>
                        <E2EDP15_TXT01>
                                <xsl:value-of select="'CUST'" />
                        </E2EDP15_TXT01>
                        <E2EDP15_TXT02></E2EDP15_TXT02>
                        <E2EDP15_TXT03></E2EDP15_TXT03>
                        <E2EDP15_TXT04></E2EDP15_TXT04>
                        <E2EDP15_TXT05></E2EDP15_TXT05>
                        <E2EDP15_TXT06></E2EDP15_TXT06>
                        <E2EDP15_TXT07></E2EDP15_TXT07>
                        <E2EDP15_TXT08></E2EDP15_TXT08>
                        <E2EDP15_TXT09></E2EDP15_TXT09>
                        <E2EDP15_TXT10></E2EDP15_TXT10>
                        <E2EDP15_TXT11></E2EDP15_TXT11>
                        <E2EDP15_TXT12></E2EDP15_TXT12>
                        <E2EDP15_TXT13></E2EDP15_TXT13>
                        <E2EDP15_TXT14></E2EDP15_TXT14>
                        <E2EDP15_TSSPRAS_ISO></E2EDP15_TSSPRAS_ISO>
                        <rest></rest>
                </E2EDP15>
        </xsl:template>

        <!--Template for 'E2EDP15' -->
        <xsl:template name="E2EDP15_01">
                <E2EDP15_01>
                        <E2EDP15_01_SEGNAM></E2EDP15_01_SEGNAM>
                        <E2EDP15_01_MANDT></E2EDP15_01_MANDT>
                        <E2EDP15_01_DOCNUM></E2EDP15_01_DOCNUM>
                        <E2EDP15_01_SEGNUM></E2EDP15_01_SEGNUM>
                        <E2EDP15_01_PSGNUM></E2EDP15_01_PSGNUM>
                        <E2EDP15_01_HLEVEL></E2EDP15_01_HLEVEL>
                        <E2EDP15_01_TDNAME>
                                <xsl:value-of select="'YADT'" />
                        </E2EDP15_01_TDNAME>
                        <E2EDP15_01_SPRAS>
                                <xsl:value-of select="'EN'" />
                        </E2EDP15_01_SPRAS>
                        <E2EDP15_01_TXTLF></E2EDP15_01_TXTLF>
                        <E2EDP15_01_TXT01>
                                <xsl:value-of select="'1210'" />
                        </E2EDP15_01_TXT01>
                        <E2EDP15_01_TXT02>
                                <xsl:value-of select="concat(../../../UNB/UNB_S002/UNB_0007_01,'*',../../../UNB/UNB_S002/UNB_0004)" />
                        </E2EDP15_01_TXT02>
                        <E2EDP15_01_TXT03>
                                <xsl:value-of select="'1310'" />
                        </E2EDP15_01_TXT03>
                        <E2EDP15_01_TXT04>
                                <xsl:choose>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'MI']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'MI']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'BY']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'BY']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP15_01_TXT04>
                        <E2EDP15_01_TXT05>
                                <xsl:value-of select="'T1PC'" />
                        </E2EDP15_01_TXT05>
                        <E2EDP15_01_TXT06>
                                <xsl:choose>
                                        <xsl:when test="../SG7/NAD_02[NAD_02_3035 = 'ST']/NAD_02_C082/NAD_02_3039 !=''">
                                                <xsl:value-of select="../SG7/NAD_02[NAD_02_3035 = 'ST']/NAD_02_C082/NAD_02_3039" />
                                        </xsl:when>
                                        <xsl:when test="../SG7/NAD_02[NAD_02_3035 = 'DP']/NAD_02_C082/NAD_02_3039 !=''">
                                                <xsl:value-of select="../SG7/NAD_02[NAD_02_3035 = 'DP']/NAD_02_C082/NAD_02_3039" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP15_01_TXT06>
                        <E2EDP15_01_TXT07>
                                <xsl:value-of select="'1420'" />
                        </E2EDP15_01_TXT07>
                        <E2EDP15_01_TXT08>
                                <xsl:choose>
                                        <xsl:when test="../SG7/NAD_02[NAD_02_3035 = 'ST']/NAD_02_C082/NAD_02_3039 !=''">
                                                <xsl:value-of select="../SG7/NAD_02[NAD_02_3035 = 'ST']/NAD_02_C082/NAD_02_3039" />
                                        </xsl:when>
                                        <xsl:when test="../SG7/NAD_02[NAD_02_3035 = 'DP']/NAD_02_C082/NAD_02_3039 !=''">
                                                <xsl:value-of select="../SG7/NAD_02[NAD_02_3035 = 'DP']/NAD_02_C082/NAD_02_3039" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP15_01_TXT08>
                        <E2EDP15_01_TXT09>
                                <xsl:value-of select="'T1UP'" />
                        </E2EDP15_01_TXT09>
                        <E2EDP15_01_TXT10>
                                <xsl:value-of select="LOC_02[LOC_02_3227 = '11']/LOC_02_C517/LOC_02_3225" />
                        </E2EDP15_01_TXT10>
                        <E2EDP15_01_TXT11>
                                <xsl:value-of select="'T1SC'" />
                        </E2EDP15_01_TXT11>
                        <E2EDP15_01_TXT12>
                                <xsl:choose>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                        <xsl:when test="../../SG2/NAD_01[NAD_01_3035 = 'SE']/NAD_01_C082/NAD_01_3039 != ''">
                                                <xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SE']/NAD_01_C082/NAD_01_3039" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP15_01_TXT12>
                        <E2EDP15_01_TXT13></E2EDP15_01_TXT13>
                        <E2EDP15_01_TXT14></E2EDP15_01_TXT14>
                        <E2EDP15_01_TSSPRAS_ISO></E2EDP15_01_TSSPRAS_ISO>
                        <rest></rest>
                </E2EDP15_01>
        </xsl:template>

        <!--Template for 'E2EDP15_02' -->
        <xsl:template name="E2EDP15_02">
                <E2EDP15_02>
                        <E2EDP15_02_SEGNAM></E2EDP15_02_SEGNAM>
                        <E2EDP15_02_MANDT></E2EDP15_02_MANDT>
                        <E2EDP15_02_DOCNUM></E2EDP15_02_DOCNUM>
                        <E2EDP15_02_SEGNUM></E2EDP15_02_SEGNUM>
                        <E2EDP15_02_PSGNUM></E2EDP15_02_PSGNUM>
                        <E2EDP15_02_HLEVEL></E2EDP15_02_HLEVEL>
                        <E2EDP15_02_TDNAME>
                                <xsl:value-of select="'YADT'" />
                        </E2EDP15_02_TDNAME>
                        <E2EDP15_02_SPRAS>
                                <xsl:value-of select="'EN'" />
                        </E2EDP15_02_SPRAS>
                        <E2EDP15_02_TXTLF></E2EDP15_02_TXTLF>
                        <E2EDP15_02_TXT01>
                                <xsl:value-of select="'1300'" />
                        </E2EDP15_02_TXT01>
                        <E2EDP15_02_TXT02>
                                <xsl:value-of select="SG15/QTY_01/QTY_01_C186[QTY_01_6063 = '48']/QTY_01_6060" />
                        </E2EDP15_02_TXT02>
                        <E2EDP15_02_TXT03>
                                <xsl:value-of select="'T1CM'" />
                        </E2EDP15_02_TXT03>
                        <E2EDP15_02_TXT04>
                                <xsl:choose>
                                        <xsl:when test="LIN/LIN_C212[LIN_7143 = 'IN']/LIN_7140 != ''">
                                                <xsl:value-of select="LIN/LIN_C212[LIN_7143 = 'IN']/LIN_7140" />
                                        </xsl:when>
                                        <xsl:when test="LIN/LIN_C212[LIN_7143 = 'BP']/LIN_7140 != ''">
                                                <xsl:value-of select="LIN/LIN_C212[LIN_7143 = 'BP']/LIN_7140" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP15_02_TXT04>
                        <E2EDP15_02_TXT05>
                                <xsl:value-of select="'1200'" />
                        </E2EDP15_02_TXT05>
                        <E2EDP15_02_TXT06>
                                <xsl:value-of select="IMD/IMD_C273/IMD_7008_01" />
                        </E2EDP15_02_TXT06>
                        <E2EDP15_02_TXT07></E2EDP15_02_TXT07>
                        <E2EDP15_02_TXT08></E2EDP15_02_TXT08>
                        <E2EDP15_02_TXT09></E2EDP15_02_TXT09>
                        <E2EDP15_02_TXT10></E2EDP15_02_TXT10>
                        <E2EDP15_02_TXT11></E2EDP15_02_TXT11>
                        <E2EDP15_02_TXT12></E2EDP15_02_TXT12>
                        <E2EDP15_02_TXT13></E2EDP15_02_TXT13>
                        <E2EDP15_02_TXT14></E2EDP15_02_TXT14>
                        <E2EDP15_02_TSSPRAS_ISO></E2EDP15_02_TSSPRAS_ISO>
                        <rest></rest>
                </E2EDP15_02>
        </xsl:template>

        <!--Template for 'E2EDP15_03' -->
        <xsl:template name="E2EDP15_03">
                <E2EDP15_03>
                        <E2EDP15_03_SEGNAM></E2EDP15_03_SEGNAM>
                        <E2EDP15_03_MANDT></E2EDP15_03_MANDT>
                        <E2EDP15_03_DOCNUM></E2EDP15_03_DOCNUM>
                        <E2EDP15_03_SEGNUM></E2EDP15_03_SEGNUM>
                        <E2EDP15_03_PSGNUM></E2EDP15_03_PSGNUM>
                        <E2EDP15_03_HLEVEL></E2EDP15_03_HLEVEL>
                        <E2EDP15_03_TDNAME>
                                <xsl:value-of select="'YADT'" />
                        </E2EDP15_03_TDNAME>
                        <E2EDP15_03_SPRAS>
                                <xsl:value-of select="'EN'" />
                        </E2EDP15_03_SPRAS>
                        <E2EDP15_03_TXTLF></E2EDP15_03_TXTLF>
                        <E2EDP15_03_TXT01>
                                <xsl:value-of select="'T1ON'" />
                        </E2EDP15_03_TXT01>
                        <E2EDP15_03_TXT02>
                                <xsl:value-of select="SG13/RFF_04/RFF_04_C506[RFF_04_1153 = 'ON']/RFF_04_1154"/>
                        </E2EDP15_03_TXT02>
                        <E2EDP15_03_TXT03>
                                <xsl:value-of select="'1180'" />
                        </E2EDP15_03_TXT03>
                        <E2EDP15_03_TXT04>
                                <xsl:choose>
                                        <xsl:when test="LIN/LIN_1082_01 != ''">
                                                <xsl:value-of select="LIN/LIN_1082_01"/>
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="'1'" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDP15_03_TXT04>
                        <E2EDP15_03_TXT05>
                                <xsl:value-of select="'T1EC'" />
                        </E2EDP15_03_TXT05>
                        <E2EDP15_03_TXT06>
                                <xsl:choose>
                                        <xsl:when test="PIA/PIA_C212[PIA_7143 = 'EC']/PIA_7140 != ''">
                                                <xsl:value-of select="PIA/PIA_C212[PIA_7143 = 'EC']/PIA_7140" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP15_03_TXT06>
                        <E2EDP15_03_TXT07></E2EDP15_03_TXT07>
                        <E2EDP15_03_TXT08></E2EDP15_03_TXT08>
                        <E2EDP15_03_TXT09></E2EDP15_03_TXT09>
                        <E2EDP15_03_TXT10></E2EDP15_03_TXT10>
                        <E2EDP15_03_TXT11></E2EDP15_03_TXT11>
                        <E2EDP15_03_TXT12></E2EDP15_03_TXT12>
                        <E2EDP15_03_TXT13></E2EDP15_03_TXT13>
                        <E2EDP15_03_TXT14></E2EDP15_03_TXT14>
                        <E2EDP15_03_TSSPRAS_ISO></E2EDP15_03_TSSPRAS_ISO>
                        <rest></rest>
                </E2EDP15_03>
        </xsl:template>

        <!--Template for 'E2EDP15_04' -->
        <xsl:template name="E2EDP15_04">
                <E2EDP15_04>
                        <E2EDP15_04_SEGNAM></E2EDP15_04_SEGNAM>
                        <E2EDP15_04_MANDT></E2EDP15_04_MANDT>
                        <E2EDP15_04_DOCNUM></E2EDP15_04_DOCNUM>
                        <E2EDP15_04_SEGNUM></E2EDP15_04_SEGNUM>
                        <E2EDP15_04_PSGNUM></E2EDP15_04_PSGNUM>
                        <E2EDP15_04_HLEVEL></E2EDP15_04_HLEVEL>
                        <E2EDP15_04_TDNAME>
                                <xsl:value-of select="'YECN'" />
                        </E2EDP15_04_TDNAME>
                        <E2EDP15_04_SPRAS>
                                <xsl:value-of select="'EN'" />
                        </E2EDP15_04_SPRAS>
                        <E2EDP15_04_TXTLF></E2EDP15_04_TXTLF>
                        <E2EDP15_04_TXT01>
                                <xsl:choose>
                                        <xsl:when test="PIA/PIA_C212[PIA_7143 = 'EC']/PIA_7140 != ''">  =
                                                <xsl:value-of select="PIA/PIA_C212[PIA_7143 = 'EC']/PIA_7140" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP15_04_TXT01>
                        <E2EDP15_04_TXT02></E2EDP15_04_TXT02>
                        <E2EDP15_04_TXT03></E2EDP15_04_TXT03>
                        <E2EDP15_04_TXT04></E2EDP15_04_TXT04>
                        <E2EDP15_04_TXT05></E2EDP15_04_TXT05>
                        <E2EDP15_04_TXT06></E2EDP15_04_TXT06>
                        <E2EDP15_04_TXT07></E2EDP15_04_TXT07>
                        <E2EDP15_04_TXT08></E2EDP15_04_TXT08>
                        <E2EDP15_04_TXT09></E2EDP15_04_TXT09>
                        <E2EDP15_04_TXT10></E2EDP15_04_TXT10>
                        <E2EDP15_04_TXT11></E2EDP15_04_TXT11>
                        <E2EDP15_04_TXT12></E2EDP15_04_TXT12>
                        <E2EDP15_04_TXT13></E2EDP15_04_TXT13>
                        <E2EDP15_04_TXT14></E2EDP15_04_TXT14>
                        <E2EDP15_04_TSSPRAS_ISO></E2EDP15_04_TSSPRAS_ISO>
                        <rest></rest>
                </E2EDP15_04>
        </xsl:template>

        <!--Template for 'E2EDP15_05' -->
        <xsl:template name="E2EDP15_05">
                <E2EDP15_05>
                        <E2EDP15_05_SEGNAM></E2EDP15_05_SEGNAM>
                        <E2EDP15_05_MANDT></E2EDP15_05_MANDT>
                        <E2EDP15_05_DOCNUM></E2EDP15_05_DOCNUM>
                        <E2EDP15_05_SEGNUM></E2EDP15_05_SEGNUM>
                        <E2EDP15_05_PSGNUM></E2EDP15_05_PSGNUM>
                        <E2EDP15_05_HLEVEL></E2EDP15_05_HLEVEL>
                        <E2EDP15_05_TDNAME>
                                <xsl:value-of select="'YC10'" />
                        </E2EDP15_05_TDNAME>
                        <E2EDP15_05_SPRAS>
                                <xsl:value-of select="'EN'" />
                        </E2EDP15_05_SPRAS>
                        <E2EDP15_05_TXTLF></E2EDP15_05_TXTLF>
                        <E2EDP15_05_TXT01>
                                <xsl:value-of select="SG15/QTY_01/QTY_01_C186[QTY_01_6063 = '48']/QTY_01_6060" />
                        </E2EDP15_05_TXT01>
                        <E2EDP15_05_TXT02></E2EDP15_05_TXT02>
                        <E2EDP15_05_TXT03></E2EDP15_05_TXT03>
                        <E2EDP15_05_TXT04></E2EDP15_05_TXT04>
                        <E2EDP15_05_TXT05></E2EDP15_05_TXT05>
                        <E2EDP15_05_TXT06></E2EDP15_05_TXT06>
                        <E2EDP15_05_TXT07></E2EDP15_05_TXT07>
                        <E2EDP15_05_TXT08></E2EDP15_05_TXT08>
                        <E2EDP15_05_TXT09></E2EDP15_05_TXT09>
                        <E2EDP15_05_TXT10></E2EDP15_05_TXT10>
                        <E2EDP15_05_TXT11></E2EDP15_05_TXT11>
                        <E2EDP15_05_TXT12></E2EDP15_05_TXT12>
                        <E2EDP15_05_TXT13></E2EDP15_05_TXT13>
                        <E2EDP15_05_TXT14></E2EDP15_05_TXT14>
                        <E2EDP15_05_TSSPRAS_ISO></E2EDP15_05_TSSPRAS_ISO>
                        <rest></rest>
                </E2EDP15_05>
        </xsl:template>

        <!--Template for 'E2EDP16' -->
        <xsl:template name="E2EDP16">
                <E2EDP16>
                        <E2EDP16_SEGNAM></E2EDP16_SEGNAM>
                        <E2EDP16_MANDT></E2EDP16_MANDT>
                        <E2EDP16_DOCNUM></E2EDP16_DOCNUM>
                        <E2EDP16_SEGNUM></E2EDP16_SEGNUM>
                        <E2EDP16_PSGNUM></E2EDP16_PSGNUM>
                        <E2EDP16_HLEVEL></E2EDP16_HLEVEL>
                        <E2EDP16_ETTYP>
                                <xsl:choose>
                                        <xsl:when test="SCC_01/SCC_01_4017 = 1 or SCC_01/SCC_01_4017 = 4">
                                                <xsl:value-of select="SCC_01/SCC_01_4017" />
                                        </xsl:when>
                                        <xsl:when test="SG18_1/QTY_02/QTY_02_C186/QTY_02_6063 = '83'">
                                                <xsl:value-of select="'1'" />
                                        </xsl:when>
                                        <xsl:when test="SG18_4/QTY_02/QTY_02_C186/QTY_02_6063 = '83'">
                                                <xsl:value-of select="'1'" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="''" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDP16_ETTYP>
                        <E2EDP16_PRGRS>
                                <xsl:choose>
                                        <xsl:when test="SCC_01/SCC_01_C329/SCC_01_2013 = 'Y'">
                                                <xsl:value-of select="'D'" />
                                        </xsl:when>
                                        <xsl:when test="SG18_1/QTY_02/QTY_02_C186/QTY_02_6063 = '83'">
                                                <xsl:value-of select="'D'" />
                                        </xsl:when>
                                        <xsl:when test="SG18_4/QTY_02/QTY_02_C186/QTY_02_6063 = '83'">
                                                <xsl:value-of select="'D'" />
                                        </xsl:when>
                                        <xsl:otherwise>
                                                <xsl:value-of select="SCC_01/SCC_01_C329/SCC_01_2013" />
                                        </xsl:otherwise>
                                </xsl:choose>
                        </E2EDP16_PRGRS>
                        <E2EDP16_EDATUV>
                                <xsl:choose>
                                        <xsl:when test="SG18_1/QTY_02/QTY_02_C186/QTY_02_6063 = '83'">
                                                <xsl:value-of select="format-date(xs:date(replace(../../../../DTM_01_gr/DTM_01_137/DTM_01_137_C507[DTM_01_137_2005 = '137']/DTM_01_137_2380,'(\d{4})(\d{2})(\d{2})','$1-$2-$3')) - xs:dayTimeDuration('P1D'), '[Y0001][M01][D01]')" />
                                        </xsl:when>
                                        <xsl:when test="SG18_4/QTY_02/QTY_02_C186/QTY_02_6063 = '83'">
                                                <xsl:value-of select="format-date(xs:date(replace(../../../../DTM_01_gr/DTM_01_137/DTM_01_137_C507[DTM_01_137_2005 = '137']/DTM_01_137_2380,'(\d{4})(\d{2})(\d{2})','$1-$2-$3')) - xs:dayTimeDuration('P1D'), '[Y0001][M01][D01]')" />
                                        </xsl:when>
                                        <xsl:when test="SG18_1[QTY_02/QTY_02_C186/QTY_02_6063 != '83']/DTM_13/DTM_13_C507[DTM_13_2005 = '2']/DTM_13_2380 != ''">
                                                <xsl:value-of select="SG18_1/DTM_13/DTM_13_C507[DTM_13_2005 = '2']/DTM_13_2380" />
                                        </xsl:when>
                                        <xsl:when test="SG18_4[QTY_02/QTY_02_C186/QTY_02_6063 != '83']/DTM_13/DTM_13_C507[DTM_13_2005 = '2']/DTM_13_2380 != ''">
                                                <xsl:value-of select="SG18_4/DTM_13/DTM_13_C507[DTM_13_2005 = '2']/DTM_13_2380" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP16_EDATUV>
                        <E2EDP16_EZEIT></E2EDP16_EZEIT>
                        <E2EDP16_EDATUB>
                                <xsl:choose>
                                        <xsl:when test="SG18_1/QTY_02/QTY_02_C186/QTY_02_6063 = '83'">
                                                <xsl:value-of select="format-date(xs:date(replace(../../../../DTM_01_gr/DTM_01_137/DTM_01_137_C507[DTM_01_137_2005 = '137']/DTM_01_137_2380,'(\d{4})(\d{2})(\d{2})','$1-$2-$3')) - xs:dayTimeDuration('P1D'), '[Y0001][M01][D01]')" />
                                        </xsl:when>
                                        <xsl:when test="SG18_4/QTY_02/QTY_02_C186/QTY_02_6063 = '83'">
                                                <xsl:value-of select="format-date(xs:date(replace(../../../../DTM_01_gr/DTM_01_137/DTM_01_137_C507[DTM_01_137_2005 = '137']/DTM_01_137_2380,'(\d{4})(\d{2})(\d{2})','$1-$2-$3')) - xs:dayTimeDuration('P1D'), '[Y0001][M01][D01]')" />
                                        </xsl:when>
                                        <xsl:when test="SG18_1[QTY_02/QTY_02_C186/QTY_02_6063 != '83']/DTM_13/DTM_13_C507[DTM_13_2005 = '2']/DTM_13_2380 != ''">
                                                <xsl:value-of select="SG18_1/DTM_13/DTM_13_C507[DTM_13_2005 = '2']/DTM_13_2380" />
                                        </xsl:when>
                                        <xsl:when test="SG18_4[QTY_02/QTY_02_C186/QTY_02_6063 != '83']/DTM_13/DTM_13_C507[DTM_13_2005 = '2']/DTM_13_2380 != ''">
                                                <xsl:value-of select="SG18_4/DTM_13/DTM_13_C507[DTM_13_2005 = '2']/DTM_13_2380" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP16_EDATUB>
                        <E2EDP16_ETVTF></E2EDP16_ETVTF>
                        <E2EDP16_WMENG>
                                <xsl:choose>
                                        <xsl:when test="SG18_1/QTY_02/QTY_02_C186[QTY_02_6063 = '1' or QTY_02_6063 = '83']/QTY_02_6060 != ''">
                                                <xsl:value-of select="SG18_1/QTY_02/QTY_02_C186[QTY_02_6063 = '1' or QTY_02_6063 = '83']/QTY_02_6060" />
                                        </xsl:when>
                                        <xsl:when test="SG18_4/QTY_02/QTY_02_C186[QTY_02_6063 = '1' or QTY_02_6063 = '83']/QTY_02_6060 != ''">
                                                <xsl:value-of select="SG18_4/QTY_02/QTY_02_C186[QTY_02_6063 = '1' or QTY_02_6063 = '83']/QTY_02_6060" />
                                        </xsl:when>
                                        <xsl:when test="SG18_1/QTY_02/QTY_02_C186[QTY_02_6063 = '131' or QTY_02_6063 = '83']/QTY_02_6060 != ''">
                                                <xsl:value-of select="SG18_1/QTY_02/QTY_02_C186[QTY_02_6063 = '131' or QTY_02_6063 = '83']/QTY_02_6060" />
                                        </xsl:when>
                                        <xsl:when test="SG18_4/QTY_02/QTY_02_C186[QTY_02_6063 = '131' or QTY_02_6063 = '83']/QTY_02_6060 != ''">
                                                <xsl:value-of select="SG18_4/QTY_02/QTY_02_C186[QTY_02_6063 = '131' or QTY_02_6063 = '83']/QTY_02_6060" />
                                        </xsl:when>
                                </xsl:choose>
                        </E2EDP16_WMENG>
                        <E2EDP16_FZABR></E2EDP16_FZABR>
                        <E2EDP16_BSTAS></E2EDP16_BSTAS>
                        <E2EDP16_WDATUV></E2EDP16_WDATUV>
                        <E2EDP16_WZEIT></E2EDP16_WZEIT>
                        <E2EDP16_WDATUB></E2EDP16_WDATUB>
                        <E2EDP16_LDATUV></E2EDP16_LDATUV>
                        <E2EDP16_LZEIT></E2EDP16_LZEIT>
                        <E2EDP16_LDATUB></E2EDP16_LDATUB>
                        <E2EDP16_BSTGRU></E2EDP16_BSTGRU>
                        <E2EDP16_TRANSPTP></E2EDP16_TRANSPTP>
                        <E2EDP16_TRANSPDESC></E2EDP16_TRANSPDESC>
                        <rest></rest>
                </E2EDP16>
        </xsl:template>
</xsl:stylesheet>
