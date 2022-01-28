<?xml version="1.0" encoding="UTF-8"?>
<!--
#
# $Header:$
#
#	Mapping: ANSI X12 862 3060 to IDOC_DELFOR02 - Kwangjin Plex
#   source description: oena_ansix12_862_4010.fde	
#   target description: oena_idoc_delfor02_tfx.fde
#
# $Log:$
# 
# date		| user							| comment
# ..............................................................................................
# 20180412	| Anjana   RBEI/BSF3		| initial mapping creation
#

they have a forecast number, but not a PO number????
why is E2EDKA1_PARTN for AG set from SU and not BY?
not sure about K11 using from COM because CTA 3412 does not exist
not sure about concat(IMD_02/IMD_02_7008_01, IMD_02/IMD_02_7383)
IMD+F++:::REAR PMA + PDC  SENSOR MELBOURNE RED WA75'




-->
<!--Stylesheet was automatically generated on Mar 14, 2018 8:01:26 PM-->
<xsl:stylesheet version="2.0"
				xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
				xmlns:xs="http://www.w3.org/2001/XMLSchema"
				xmlns:fn="http://www.w3.org/2005/xpath-functions"
				xmlns:err="http://www.w3.org/2005/xqt-errors"
				xmlns:XmlTransform="java:de.axway.converter.xml.xsltextensions.XsltExtensions"
					exclude-result-prefixes="XmlTransform"
				xmlns:saxon="http://saxon.sf.net/"
				xmlns:bosch="http://www.bosch.com"
				xmlns:exslt="http://exslt.org/common"
				extension-element-prefixes="saxon xs fn err bosch XmlTransform exslt">
				
	<!-- ############################################################### -->
	<!--  Import                                                         -->
	<!-- ############################################################### -->

	<xsl:import href="RBNA_FUNCTIONS_XSLT_UTILS.xsl" />
	<xsl:import href="RBNA_TEMPLATE_FUNCTIONS.xsl" />
    
    <!-- ############################################################### -->
	<!--  Parameters and Variables                                       -->
	<!-- ############################################################### -->
    
  <xsl:param name="ACTISEDI" select="''"/>
	<xsl:param name="PPATH" select="''"/>
	<xsl:param name="REF_NR" select="'12345'" saxon:assignable="yes"/>
	<xsl:param name="stepfile" select="''" />
	<xsl:param name="MANDT" select="''"/>
	<xsl:param name="CIMTYP" select="''"/>
	<xsl:param name="IDOCTYP" select="''"/>
	<xsl:param name="MESTYP" select="''"/>
	<xsl:param name="MESCOD" select="''"/>
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
	<xsl:param name="JOB_ID" select="''"/>
	<xsl:param name="JOBID" select="''"/>
	<xsl:param name="OBJECT_ID" select="''"/>
	
	<!--  date, time from XSLT 2.0 functions-->
	<xsl:variable name="gv_docdate" select="format-date(current-date(),'[Y0001][M01][D01]')"/>
	<xsl:variable name="gv_doctime" select="format-time(current-time(),'[H01][m01][s01]')"/>
	<xsl:variable name="gv_edidate" select="format-date(current-date(),'[Y01][M01][D01]')" />
	<xsl:variable name="gv_editime" select="format-time(current-time(),'[H01][m01]')" />
	
	<!-- saxon:assignable variables, need for counters -->
	<xsl:variable name="gv_lin_counter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_unh_counter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_message_reference_number" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_segment_counter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_st_counter" select="0" saxon:assignable="yes"/>	
	<xsl:variable name="gv_E2EDK09_LABNK_counter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_fstcounter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_LIN_235" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_LIN_234" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_VendorsPartNumber" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_BuyersPartNumber" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_BatchNumber" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_POLineItemNumber" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_RevisionNumber" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_EngineeringChangeNumber" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_ProductDescription" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_LIN_PO_Number" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_ConfigurationData" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_LINPL" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_AG_PARTN" select="''" saxon:assignable = "yes" />
	<xsl:variable name="gv_WK_PARTN" select="''" saxon:assignable = "yes" />
	
	<!-- SNDPRN /-->
	<xsl:variable name="gv_SNDPRN" select="XmlTransform:getPropertyActisedi('/lib/convert/tftab/OENA_CODETAB_INBOUND_PARTNER',concat(/M_DELFOR_D96A/S_UNB/UNB_0004,/M_DELFOR_D96A/S_UNB/UNB_0010,'#SNDPRN'))" saxon:assignable="yes"/>
	
	<!-- Mapping Name /-->
	<xsl:variable name="gv_mapping" select="'oena_ansix12_862_4010_kwangjin_plex.oena_idoc_delfor02.xsl'" saxon:assignable="yes"/>
	
	<!-- output variable -->
	<xsl:param name="resultfile" select="''" />

	<!-- ############################################################### -->
	<!--  START Mapping                                                  -->
	<!-- ############################################################### -->
	
	<xsl:output method="xml" indent="yes" encoding="UTF-8"/>
	
	<!-- ######################################## Root ######################################## -->

	<!--Template for root element -->
	<xsl:template match="/">
	
		<!-- write stepfile name to log file -->
		<xsl:message terminate="no">
			<xsl:text>STEPFILE: </xsl:text>
			<xsl:value-of select="$stepfile" />
		</xsl:message>
		
		<!-- write mapping name to log file -->
		<xsl:message terminate="no">
			<xsl:text>XSL: </xsl:text>
			<xsl:value-of select="$gv_mapping" />
		</xsl:message>
		<xsl:message select="concat(/M_DELFOR_D96A/S_UNB/UNB_0004,/M_DELFOR_D96A/S_UNB/UNB_0010,'#SNDPRN')"></xsl:message>
		<xsl:choose>
			<xsl:when test="$JOBID != ''">
				<saxon:assign name="REF_NR" select="$JOBID"/>
			</xsl:when>
			<xsl:when test="$JOB_ID != ''">
				<saxon:assign name="REF_NR" select="$JOB_ID"/>
			</xsl:when>
			<xsl:when test="$OBJECT_ID != ''">
				<saxon:assign name="REF_NR" select="$OBJECT_ID"/>
			</xsl:when>
			<xsl:otherwise>
				<saxon:assign name="REF_NR" select="$REF_NR"/>
			</xsl:otherwise>
		</xsl:choose>

		<xsl:call-template name="IDoc" />
		
		<xsl:choose>
			<xsl:when test="$gv_SNDPRN"/>
			<xsl:otherwise>
				<xsl:message terminate="yes">
					<xsl:text>ERROR: no value found in the table OENA_CODETAB_INBOUND_PARTNER for </xsl:text>
					<xsl:text>EDICode : </xsl:text><xsl:value-of select="concat(/X12_4010_862/S_GS/GS_142,/X12_4010_862/S_GS/GS_124,'#SNDPRN')" />
				</xsl:message>
			</xsl:otherwise>
		</xsl:choose>

	</xsl:template>

	<!--Template for 'IDoc' -->
	<xsl:template name="IDoc">
		<IDoc>
			<xsl:for-each select="/M_DELFOR_D96A/SG0/SG4/SG8">
				<xsl:call-template name="G00" />
			</xsl:for-each>
		</IDoc>
	</xsl:template>

	<!--Template for 'G00' -->
	<xsl:template name="G00">
		<G00>
			<xsl:call-template name="EDI_DC40" />
			<xsl:call-template name="G01" />
		</G00>
	</xsl:template>

	<!--Template for 'EDI_DC40' -->
	<xsl:template name="EDI_DC40">
		<EDI_DC40>
			<EDI_DC40_TABNAM>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_TABNAM>
			<EDI_DC40_MANDT>
				<xsl:value-of select="$MANDT" />
			</EDI_DC40_MANDT>
			<EDI_DC40_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_DOCNUM>
			<EDI_DC40_DOCREL>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_DOCREL>
			<EDI_DC40_STATUS>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_STATUS>
			<EDI_DC40_DIRECT>
				<xsl:value-of select="'2'" />
			</EDI_DC40_DIRECT>
			<EDI_DC40_OUTMOD>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_OUTMOD>
			<EDI_DC40_EXPRSS>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_EXPRSS>
			<EDI_DC40_TEST>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_TEST>
			<EDI_DC40_IDOCTYP>
				<xsl:value-of select="'DELFOR02'" />
			</EDI_DC40_IDOCTYP>
			<EDI_DC40_CIMTYP>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_CIMTYP>
			<EDI_DC40_MESTYP>
				<xsl:value-of select="'DELINS'" />
			</EDI_DC40_MESTYP>
			<EDI_DC40_MESCOD>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_MESCOD>
			<EDI_DC40_MESFCT>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_MESFCT>
			<EDI_DC40_STD>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_STD>
			<EDI_DC40_STDVRS>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_STDVRS>
			<EDI_DC40_STDMES>
				<xsl:value-of select="'830'" />
			</EDI_DC40_STDMES>
			<EDI_DC40_SNDPOR>
				<xsl:value-of select="$SNDPOR" />
			</EDI_DC40_SNDPOR>
			<EDI_DC40_SNDPRT>
				<xsl:value-of select="$SNDPRT" />
			</EDI_DC40_SNDPRT>
			<EDI_DC40_SNDPFC>
				<xsl:value-of select="$SNDPFC" />
			</EDI_DC40_SNDPFC>
			<EDI_DC40_SNDPRN>
				<xsl:value-of select="$gv_SNDPRN" />
			</EDI_DC40_SNDPRN>
			<EDI_DC40_SNDSAD>
				<xsl:value-of select="'PO'" />
			</EDI_DC40_SNDSAD>
			<EDI_DC40_SNDLAD>
				<xsl:value-of select="concat(../NAD_02[NAD_02_3035 = 'ST']/NAD_02_3039,';',../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_3039,';')"/>						
			</EDI_DC40_SNDLAD>
			<EDI_DC40_RCVPOR>
				<xsl:value-of select="$RCVPOR" />
			</EDI_DC40_RCVPOR>
			<EDI_DC40_RCVPRT>
				<xsl:value-of select="$RCVPRT" />
			</EDI_DC40_RCVPRT>
			<EDI_DC40_RCVPFC>
				<xsl:value-of select="$RCVPFC" />
			</EDI_DC40_RCVPFC>
			<EDI_DC40_RCVPRN>
				<xsl:value-of select="$RCVPRN" />
			</EDI_DC40_RCVPRN>
			<EDI_DC40_RCVSAD>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_RCVSAD>
			<EDI_DC40_RCVLAD>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_RCVLAD>
			<EDI_DC40_CREDAT>
				<xsl:value-of select="$gv_docdate"/>
			</EDI_DC40_CREDAT>
			<EDI_DC40_CRETIM>
				<xsl:value-of select="$gv_doctime"/>
			</EDI_DC40_CRETIM>
			<EDI_DC40_REFINT>
				<xsl:value-of select="../../../S_UNB/UNB_0019"/>
			</EDI_DC40_REFINT>
			<EDI_DC40_REFGRP>
				<xsl:value-of select="../../UNH/UNH_0062"/>
			</EDI_DC40_REFGRP>
			<EDI_DC40_REFMES>
				<xsl:value-of select="LIN_01/LIN_01_1082_01"/>
			</EDI_DC40_REFMES>
			<EDI_DC40_ARCKEY>
				<xsl:choose>
					<xsl:when test="$JOBID != ''">
						<xsl:value-of select="concat('workflow-ID: ',$JOBID)"/>
					</xsl:when>
					<xsl:when test="$JOB_ID != ''">
						<xsl:value-of select="concat('workflow-ID: ',$JOB_ID)"/>
					</xsl:when>
					<xsl:when test="$OBJECT_ID != ''">
						<xsl:value-of select="concat('workflow-ID: ',$OBJECT_ID)"/>
					</xsl:when>
					<xsl:otherwise></xsl:otherwise>
				</xsl:choose>
			</EDI_DC40_ARCKEY>
			<EDI_DC40_SERIAL>
				<!--Please insert the transformation instruction here-->
			</EDI_DC40_SERIAL>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</EDI_DC40>
	</xsl:template>

	<!--Template for 'G01' -->
	<xsl:template name="G01">
		<G01>
			<xsl:call-template name="E2EDK09" />
			<xsl:call-template name="E2EDK10" />
			<xsl:call-template name="E2EDKA1" />
			<xsl:call-template name="E2EDKA1_AG" />
			<xsl:call-template name="E2EDKA1_WK" />
			<xsl:for-each select="../../SG2/SG3/COM_01">
				<xsl:call-template name="E2EDK11" />
			</xsl:for-each>
			<xsl:call-template name="G02" />
		</G01>
	</xsl:template>

	<!--Template for 'E2EDK09' -->
	<xsl:template name="E2EDK09">
		<E2EDK09>
			<E2EDK09_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_SEGNAM>
			<E2EDK09_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_MANDT>
			<E2EDK09_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_DOCNUM>
			<E2EDK09_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_SEGNUM>
			<E2EDK09_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_PSGNUM>
			<E2EDK09_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_HLEVEL>
			<E2EDK09_VTRNR>
				<xsl:value-of select="../../BGM_01/BGM_01_1004"/>
			</E2EDK09_VTRNR>
			<E2EDK09_BSTDK>
				<xsl:value-of select="substring(../../DTM_01[DTM_01_2005 = '137']/DTM_01_2380, 1, 8)"/>
			</E2EDK09_BSTDK>
			<E2EDK09_LABNK>
				<xsl:value-of select="../../BGM_01/BGM_01_1004"/>
			</E2EDK09_LABNK>
			<E2EDK09_ZEICH>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_ZEICH>
			<E2EDK09_BSTZD>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_BSTZD>
			<E2EDK09_ABRVW>
				<xsl:value-of select="'S'" />
			</E2EDK09_ABRVW>
			<E2EDK09_KSTAT>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_KSTAT>
			<E2EDK09_KTEXT>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_KTEXT>
			<E2EDK09_ABNRA>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_ABNRA>
			<E2EDK09_ABNRD>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_ABNRD>
			<E2EDK09_KTEXT_V>
				<xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_3036_01"/>
			</E2EDK09_KTEXT_V>
			<E2EDK09_USR01>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_USR01>
			<E2EDK09_USR02>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_USR02>
			<E2EDK09_USR03>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_USR03>
			<E2EDK09_USR04>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_USR04>
			<E2EDK09_USR05>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_USR05>
			<E2EDK09_CYEFZ>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_CYEFZ>
			<E2EDK09_CYDAT>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_CYDAT>
			<E2EDK09_MFLAUF>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_MFLAUF>
			<E2EDK09_MFEIN>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_MFEIN>
			<E2EDK09_FFLAUF>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_FFLAUF>
			<E2EDK09_FFEIN>
				<!--Please insert the transformation instruction here-->
			</E2EDK09_FFEIN>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDK09>
	</xsl:template>

	<!--Template for 'E2EDK10' -->
	<xsl:template name="E2EDK10">
		<E2EDK10>
			<E2EDK10_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDK10_SEGNAM>
			<E2EDK10_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDK10_MANDT>
			<E2EDK10_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK10_DOCNUM>
			<E2EDK10_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK10_SEGNUM>
			<E2EDK10_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK10_PSGNUM>
			<E2EDK10_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDK10_HLEVEL>
			<E2EDK10_DUMMY>
				<!--Please insert the transformation instruction here-->
			</E2EDK10_DUMMY>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDK10>
	</xsl:template>

	<!--Template for 'E2EDKA1' -->
	<xsl:template name="E2EDKA1">
		<E2EDKA1>
			<E2EDKA1_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SEGNAM>
			<E2EDKA1_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_MANDT>
			<E2EDKA1_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_DOCNUM>
			<E2EDKA1_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SEGNUM>
			<E2EDKA1_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSGNUM>
			<E2EDKA1_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_HLEVEL>
			<E2EDKA1_PARVW>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PARVW>
			<E2EDKA1_PARTN>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PARTN>
			<E2EDKA1_LIFNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_LIFNR>
			<E2EDKA1_NAME1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME1>
			<E2EDKA1_NAME2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME2>
			<E2EDKA1_NAME3>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME3>
			<E2EDKA1_NAME4>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME4>
			<E2EDKA1_STRAS>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STRAS>
			<E2EDKA1_STRS2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STRS2>
			<E2EDKA1_PFACH>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PFACH>
			<E2EDKA1_ORT01>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORT01>
			<E2EDKA1_COUNC>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_COUNC>
			<E2EDKA1_PSTLZ>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSTLZ>
			<E2EDKA1_PSTL2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSTL2>
			<E2EDKA1_LAND1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_LAND1>
			<E2EDKA1_ABLAD>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ABLAD>
			<E2EDKA1_PERNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PERNR>
			<E2EDKA1_PARNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PARNR>
			<E2EDKA1_TELF1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELF1>
			<E2EDKA1_TELF2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELF2>
			<E2EDKA1_TELBX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELBX>
			<E2EDKA1_TELFX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELFX>
			<E2EDKA1_TELTX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELTX>
			<E2EDKA1_TELX1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELX1>
			<E2EDKA1_SPRAS>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SPRAS>
			<E2EDKA1_ANRED>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ANRED>
			<E2EDKA1_ORT02>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORT02>
			<E2EDKA1_HAUSN>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_HAUSN>
			<E2EDKA1_STOCK>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STOCK>
			<E2EDKA1_REGIO>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_REGIO>
			<E2EDKA1_PARGE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PARGE>
			<E2EDKA1_ISOAL>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ISOAL>
			<E2EDKA1_ISONU>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ISONU>
			<E2EDKA1_FCODE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_FCODE>
			<E2EDKA1_IHREZ>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_IHREZ>
			<E2EDKA1_BNAME>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_BNAME>
			<E2EDKA1_PAORG>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PAORG>
			<E2EDKA1_ORGTX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORGTX>
			<E2EDKA1_PAGRU>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PAGRU>
			<E2EDKA1_KNREF>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_KNREF>
			<E2EDKA1_ILNNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ILNNR>
			<E2EDKA1_PFORT>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PFORT>
			<E2EDKA1_SPRAS_ISO>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SPRAS_ISO>
			<E2EDKA1_TITLE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TITLE>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDKA1>
	</xsl:template>

	<!--Template for 'E2EDKA1_AG' -->
	<xsl:template name="E2EDKA1_AG">
		<E2EDKA1>
			<E2EDKA1_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SEGNAM>
			<E2EDKA1_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_MANDT>
			<E2EDKA1_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_DOCNUM>
			<E2EDKA1_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SEGNUM>
			<E2EDKA1_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSGNUM>
			<E2EDKA1_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_HLEVEL>
			<E2EDKA1_PARVW>
				<xsl:value-of select="'AG'"></xsl:value-of>
			</E2EDKA1_PARVW>
			<E2EDKA1_PARTN>
				<xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_3039"/>
				<saxon:assign name="gv_AG_PARTN" select="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_3039"/>
			</E2EDKA1_PARTN>
			<E2EDKA1_LIFNR>
				<xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_3039"/>
			</E2EDKA1_LIFNR>
			<E2EDKA1_NAME1>
				<xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'BY']/NAD_01_3036_01"/>
			</E2EDKA1_NAME1>
			<E2EDKA1_NAME2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME2>
			<E2EDKA1_NAME3>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME3>
			<E2EDKA1_NAME4>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME4>
			<E2EDKA1_STRAS>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STRAS>
			<E2EDKA1_STRS2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STRS2>
			<E2EDKA1_PFACH>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PFACH>
			<E2EDKA1_ORT01>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORT01>
			<E2EDKA1_COUNC>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_COUNC>
			<E2EDKA1_PSTLZ>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSTLZ>
			<E2EDKA1_PSTL2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSTL2>
			<E2EDKA1_LAND1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_LAND1>
			<E2EDKA1_ABLAD>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ABLAD>
			<E2EDKA1_PERNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PERNR>
			<E2EDKA1_PARNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PARNR>
			<E2EDKA1_TELF1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELF1>
			<E2EDKA1_TELF2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELF2>
			<E2EDKA1_TELBX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELBX>
			<E2EDKA1_TELFX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELFX>
			<E2EDKA1_TELTX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELTX>
			<E2EDKA1_TELX1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELX1>
			<E2EDKA1_SPRAS>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SPRAS>
			<E2EDKA1_ANRED>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ANRED>
			<E2EDKA1_ORT02>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORT02>
			<E2EDKA1_HAUSN>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_HAUSN>
			<E2EDKA1_STOCK>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STOCK>
			<E2EDKA1_REGIO>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_REGIO>
			<E2EDKA1_PARGE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PARGE>
			<E2EDKA1_ISOAL>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ISOAL>
			<E2EDKA1_ISONU>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ISONU>
			<E2EDKA1_FCODE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_FCODE>
			<E2EDKA1_IHREZ>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_IHREZ>
			<E2EDKA1_BNAME>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_BNAME>
			<E2EDKA1_PAORG>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PAORG>
			<E2EDKA1_ORGTX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORGTX>
			<E2EDKA1_PAGRU>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PAGRU>
			<E2EDKA1_KNREF>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_KNREF>
			<E2EDKA1_ILNNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ILNNR>
			<E2EDKA1_PFORT>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PFORT>
			<E2EDKA1_SPRAS_ISO>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SPRAS_ISO>
			<E2EDKA1_TITLE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TITLE>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDKA1>
	</xsl:template>

	<!--Template for 'E2EDKA1_WK' -->
	<xsl:template name="E2EDKA1_WK">
		<E2EDKA1>
			<E2EDKA1_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SEGNAM>
			<E2EDKA1_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_MANDT>
			<E2EDKA1_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_DOCNUM>
			<E2EDKA1_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SEGNUM>
			<E2EDKA1_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSGNUM>
			<E2EDKA1_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_HLEVEL>
			<E2EDKA1_PARVW>
				<xsl:value-of select="'WK'"></xsl:value-of>
			</E2EDKA1_PARVW>
			<E2EDKA1_PARTN>
				<xsl:value-of select="../NAD_02[NAD_02_3035 = 'ST']/NAD_02_3039"/>
				<saxon:assign name="gv_WK_PARTN" select="../NAD_02[NAD_02_3035 = 'ST']/NAD_02_3039"/>
			</E2EDKA1_PARTN>
			<E2EDKA1_LIFNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_LIFNR>
			<E2EDKA1_NAME1>
				<xsl:value-of select="../NAD_02[NAD_02_3035 = 'ST']/NAD_02_3036_01"/>
			</E2EDKA1_NAME1>
			<E2EDKA1_NAME2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME2>
			<E2EDKA1_NAME3>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME3>
			<E2EDKA1_NAME4>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_NAME4>
			<E2EDKA1_STRAS>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STRAS>
			<E2EDKA1_STRS2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STRS2>
			<E2EDKA1_PFACH>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PFACH>
			<E2EDKA1_ORT01>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORT01>
			<E2EDKA1_COUNC>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_COUNC>
			<E2EDKA1_PSTLZ>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSTLZ>
			<E2EDKA1_PSTL2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PSTL2>
			<E2EDKA1_LAND1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_LAND1>
			<E2EDKA1_ABLAD>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ABLAD>
			<E2EDKA1_PERNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PERNR>
			<E2EDKA1_PARNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PARNR>
			<E2EDKA1_TELF1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELF1>
			<E2EDKA1_TELF2>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELF2>
			<E2EDKA1_TELBX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELBX>
			<E2EDKA1_TELFX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELFX>
			<E2EDKA1_TELTX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELTX>
			<E2EDKA1_TELX1>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TELX1>
			<E2EDKA1_SPRAS>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SPRAS>
			<E2EDKA1_ANRED>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ANRED>
			<E2EDKA1_ORT02>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORT02>
			<E2EDKA1_HAUSN>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_HAUSN>
			<E2EDKA1_STOCK>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_STOCK>
			<E2EDKA1_REGIO>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_REGIO>
			<E2EDKA1_PARGE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PARGE>
			<E2EDKA1_ISOAL>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ISOAL>
			<E2EDKA1_ISONU>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ISONU>
			<E2EDKA1_FCODE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_FCODE>
			<E2EDKA1_IHREZ>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_IHREZ>
			<E2EDKA1_BNAME>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_BNAME>
			<E2EDKA1_PAORG>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PAORG>
			<E2EDKA1_ORGTX>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ORGTX>
			<E2EDKA1_PAGRU>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PAGRU>
			<E2EDKA1_KNREF>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_KNREF>
			<E2EDKA1_ILNNR>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_ILNNR>
			<E2EDKA1_PFORT>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_PFORT>
			<E2EDKA1_SPRAS_ISO>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_SPRAS_ISO>
			<E2EDKA1_TITLE>
				<!--Please insert the transformation instruction here-->
			</E2EDKA1_TITLE>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDKA1>
	</xsl:template>

	<!--Template for 'E2EDK11' -->
	<xsl:template name="E2EDK11">
		<E2EDK11>
			<E2EDK11_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_SEGNAM>
			<E2EDK11_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_MANDT>
			<E2EDK11_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_DOCNUM>
			<E2EDK11_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_SEGNUM>
			<E2EDK11_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_PSGNUM>
			<E2EDK11_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_HLEVEL>
			<E2EDK11_TDNAME>
				<xsl:value-of select="'Z830'"/>
			</E2EDK11_TDNAME>
			<E2EDK11_SPRAS>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_SPRAS>
			<E2EDK11_TXTLF>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXTLF>
			<E2EDK11_TXT01>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT01>
			<E2EDK11_TXT02>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT02>
			<E2EDK11_TXT03>
				<xsl:value-of select="COM_01_3148"/>
			</E2EDK11_TXT03>
			<E2EDK11_TXT04>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT04>
			<E2EDK11_TXT05>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT05>
			<E2EDK11_TXT06>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT06>
			<E2EDK11_TXT07>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT07>
			<E2EDK11_TXT08>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT08>
			<E2EDK11_TXT09>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT09>
			<E2EDK11_TXT10>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT10>
			<E2EDK11_TXT11>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT11>
			<E2EDK11_TXT12>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT12>
			<E2EDK11_TXT13>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT13>
			<E2EDK11_TXT14>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TXT14>
			<E2EDK11_TSSPRAS_ISO>
				<!--Please insert the transformation instruction here-->
			</E2EDK11_TSSPRAS_ISO>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDK11>
	</xsl:template>

	<!--Template for 'G02' -->
	<xsl:template name="G02">
		<G02>
			<xsl:call-template name="E2EDP10" />
			<xsl:call-template name="E2EDP11" />
			<xsl:call-template name="E2EDP36" />
			<xsl:call-template name="E2EDP14" />
			<xsl:call-template name="E2EDP15" />
			<xsl:call-template name="E2EDP15_YAQU" />
			<xsl:call-template name="E2EDP15_YADT_1260" />
			<xsl:call-template name="E2EDP15_YADT_1180" />
			<xsl:for-each select="SG12">
				<xsl:if test="QTY_01/QTY_01_6060 != 0 and not(empty(QTY_01/QTY_01_6060))">
					<xsl:call-template name="E2EDP16" />
				</xsl:if>
			</xsl:for-each>
		</G02>
	</xsl:template>

	<!--Template for 'E2EDP10' -->
	<xsl:template name="E2EDP10">
		<E2EDP10>
			<E2EDP10_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_SEGNAM>
			<E2EDP10_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_MANDT>
			<E2EDP10_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_DOCNUM>
			<E2EDP10_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_SEGNUM>
			<E2EDP10_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_PSGNUM>
			<E2EDP10_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_HLEVEL>
			<E2EDP10_IDNKD>
				<xsl:value-of select="LIN_01[LIN_01_7143 = 'BP']/LIN_01_7140"/>
			</E2EDP10_IDNKD>
			<E2EDP10_ARKTX>
				<xsl:value-of select="concat($gv_WK_PARTN, '   ', $gv_AG_PARTN)"></xsl:value-of>
			</E2EDP10_ARKTX>
			<E2EDP10_VRKME>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_VRKME>
			<E2EDP10_KWERK>
				<xsl:value-of select="../NAD_02[NAD_02_3035 = 'ST']/NAD_02_3039"/>
			</E2EDP10_KWERK>
			<E2EDP10_KLGOR>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_KLGOR>
			<E2EDP10_DFABL>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_DFABL>
			<E2EDP10_VBRST>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_VBRST>
			<E2EDP10_BELNR>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_BELNR>
			<E2EDP10_LFIMG>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_LFIMG>
			<E2EDP10_VEMNG>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_VEMNG>
			<E2EDP10_LIDTL>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_LIDTL>
			<E2EDP10_LIFST>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_LIFST>
			<E2EDP10_KRITB>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_KRITB>
			<E2EDP10_ABHOR>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_ABHOR>
			<E2EDP10_ABDAT>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_ABDAT>
			<E2EDP10_FZDIF>
				<xsl:value-of select="'0'"></xsl:value-of>
			</E2EDP10_FZDIF>
			<E2EDP10_AKUBM>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_AKUBM>
			<E2EDP10_AKUEM>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_AKUEM>
			<E2EDP10_DESRE>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_DESRE>
			<E2EDP10_ECHNO>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_ECHNO>
			<E2EDP10_ECHDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_ECHDT>
			<E2EDP10_ABFDE>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_ABFDE>
			<E2EDP10_MFADT>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_MFADT>
			<E2EDP10_ABMDE>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_ABMDE>
			<E2EDP10_FPSDN>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_FPSDN>
			<E2EDP10_LABKY>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_LABKY>
			<E2EDP10_TSTKY>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_TSTKY>
			<E2EDP10_DOKPF>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_DOKPF>
			<E2EDP10_KZAUS>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_KZAUS>
			<E2EDP10_ABRAB>
				<xsl:value-of select="../../DTM_01[DTM_01_2005 = '194']/DTM_01_2380"/>
			</E2EDP10_ABRAB>
			<E2EDP10_ABRBI>
				<xsl:value-of select="../../DTM_01[DTM_01_2005 = '206']/DTM_01_2380"/>
			</E2EDP10_ABRBI>
			<E2EDP10_SCREL>
				<xsl:value-of select="'03'"/>
			</E2EDP10_SCREL>
			<E2EDP10_NULDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_NULDT>
			<E2EDP10_TEART>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_TEART>
			<E2EDP10_SOLLFZ>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_SOLLFZ>
			<E2EDP10_SOLLDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_SOLLDT>
			<E2EDP10_IDNLF>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_IDNLF>
			<E2EDP10_BSTDK>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_BSTDK>
			<E2EDP10_LABNK>
				<xsl:value-of select="../BGM_01/BGM_01_1004"/>
			</E2EDP10_LABNK>
			<E2EDP10_ABNRA>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_ABNRA>
			<E2EDP10_ABNRD>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_ABNRD>
			<E2EDP10_VTRNR>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_VTRNR>
			<E2EDP10_MFRFZ>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_MFRFZ>
			<E2EDP10_POSEX>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_POSEX>
			<E2EDP10_IDNKD_EXTERNAL>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_IDNKD_EXTERNAL>
			<E2EDP10_IDNKD_GUID>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_IDNKD_GUID>
			<E2EDP10_IDNKD_VERSION>
				<!--Please insert the transformation instruction here-->
			</E2EDP10_IDNKD_VERSION>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP10>
	</xsl:template>

	<!--Template for 'E2EDP11' -->
	<xsl:template name="E2EDP11">
		<E2EDP11>
			<E2EDP11_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP11_SEGNAM>
			<E2EDP11_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP11_MANDT>
			<E2EDP11_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP11_DOCNUM>
			<E2EDP11_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP11_SEGNUM>
			<E2EDP11_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP11_PSGNUM>
			<E2EDP11_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP11_HLEVEL>
			<E2EDP11_DUMMY>
				<!--Please insert the transformation instruction here-->
			</E2EDP11_DUMMY>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP11>
	</xsl:template>

	<!--Template for 'E2EDP36' -->
	<xsl:template name="E2EDP36">
		<E2EDP36>
			<E2EDP36_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_SEGNAM>
			<E2EDP36_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_MANDT>
			<E2EDP36_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_DOCNUM>
			<E2EDP36_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_SEGNUM>
			<E2EDP36_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_PSGNUM>
			<E2EDP36_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_HLEVEL>
			<E2EDP36_LIFST>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_LIFST>
			<E2EDP36_BELNR>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_BELNR>
			<E2EDP36_BELDAT>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_BELDAT>
			<E2EDP36_BELTIM>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_BELTIM>
			<E2EDP36_LFIMG>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_LFIMG>
			<E2EDP36_VRKME>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_VRKME>
			<E2EDP36_VEMNG>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_VEMNG>
			<E2EDP36_LIKME>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_LIKME>
			<E2EDP36_LIDTL>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_LIDTL>
			<E2EDP36_LIDTIM>
				<!--Please insert the transformation instruction here-->
			</E2EDP36_LIDTIM>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP36>
	</xsl:template>

	<!--Template for 'E2EDP14' -->
	<xsl:template name="E2EDP14">
		<E2EDP14>
			<E2EDP14_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_SEGNAM>
			<E2EDP14_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_MANDT>
			<E2EDP14_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_DOCNUM>
			<E2EDP14_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_SEGNUM>
			<E2EDP14_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_PSGNUM>
			<E2EDP14_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_HLEVEL>
			<E2EDP14_PCKAR>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_PCKAR>
			<E2EDP14_PCKNR>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_PCKNR>
			<E2EDP14_ANZPK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_ANZPK>
			<E2EDP14_PCKSL>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_PCKSL>
			<E2EDP14_ANZAR>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_ANZAR>
			<E2EDP14_MAZAR>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_MAZAR>
			<E2EDP14_MMPCK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_MMPCK>
			<E2EDP14_LPACK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_LPACK>
			<E2EDP14_BPACK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_BPACK>
			<E2EDP14_HPACK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_HPACK>
			<E2EDP14_MGPCK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_MGPCK>
			<E2EDP14_GPACK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_GPACK>
			<E2EDP14_NPACK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_NPACK>
			<E2EDP14_VPACK>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_VPACK>
			<E2EDP14_GDATV>
				<!--Please insert the transformation instruction here-->
			</E2EDP14_GDATV>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP14>
	</xsl:template>

	<!--Template for 'E2EDP15' -->
	<xsl:template name="E2EDP15">
		<E2EDP15>
			<E2EDP15_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SEGNAM>
			<E2EDP15_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_MANDT>
			<E2EDP15_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_DOCNUM>
			<E2EDP15_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SEGNUM>
			<E2EDP15_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_PSGNUM>
			<E2EDP15_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_HLEVEL>
			<E2EDP15_TDNAME>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TDNAME>
			<E2EDP15_SPRAS>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SPRAS>
			<E2EDP15_TXTLF>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXTLF>
			<E2EDP15_TXT01>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT01>
			<E2EDP15_TXT02>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT02>
			<E2EDP15_TXT03>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT03>
			<E2EDP15_TXT04>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT04>
			<E2EDP15_TXT05>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT05>
			<E2EDP15_TXT06>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT06>
			<E2EDP15_TXT07>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT07>
			<E2EDP15_TXT08>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT08>
			<E2EDP15_TXT09>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT09>
			<E2EDP15_TXT10>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT10>
			<E2EDP15_TXT11>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT11>
			<E2EDP15_TXT12>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT12>
			<E2EDP15_TXT13>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT13>
			<E2EDP15_TXT14>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT14>
			<E2EDP15_TSSPRAS_ISO>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TSSPRAS_ISO>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP15>
	</xsl:template>

	<!--Template for 'E2EDP15_YAQU' -->
	<xsl:template name="E2EDP15_YAQU">
		<E2EDP15>
			<E2EDP15_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SEGNAM>
			<E2EDP15_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_MANDT>
			<E2EDP15_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_DOCNUM>
			<E2EDP15_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SEGNUM>
			<E2EDP15_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_PSGNUM>
			<E2EDP15_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_HLEVEL>
			<E2EDP15_TDNAME>
				<xsl:value-of select="'YAQU'"></xsl:value-of>
			</E2EDP15_TDNAME>
			<E2EDP15_SPRAS>
				<xsl:value-of select="'E'"></xsl:value-of>
			</E2EDP15_SPRAS>
			<E2EDP15_TXTLF>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXTLF>
			<E2EDP15_TXT01>
				<xsl:value-of select="'CUST'"></xsl:value-of>
			</E2EDP15_TXT01>
			<E2EDP15_TXT02>
				<xsl:value-of select="../../SG2/NAD_01[NAD_01_3035 = 'SU']/NAD_01_3039"/>
			</E2EDP15_TXT02>
			<E2EDP15_TXT03>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT03>
			<E2EDP15_TXT04>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT04>
			<E2EDP15_TXT05>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT05>
			<E2EDP15_TXT06>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT06>
			<E2EDP15_TXT07>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT07>
			<E2EDP15_TXT08>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT08>
			<E2EDP15_TXT09>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT09>
			<E2EDP15_TXT10>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT10>
			<E2EDP15_TXT11>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT11>
			<E2EDP15_TXT12>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT12>
			<E2EDP15_TXT13>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT13>
			<E2EDP15_TXT14>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT14>
			<E2EDP15_TSSPRAS_ISO>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TSSPRAS_ISO>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP15>
	</xsl:template>

	<!--Template for 'E2EDP15_YADT_1260' -->
	<xsl:template name="E2EDP15_YADT_1260">
		<E2EDP15>
			<E2EDP15_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SEGNAM>
			<E2EDP15_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_MANDT>
			<E2EDP15_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_DOCNUM>
			<E2EDP15_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SEGNUM>
			<E2EDP15_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_PSGNUM>
			<E2EDP15_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_HLEVEL>
			<E2EDP15_TDNAME>
				<xsl:value-of select="'YADT'"></xsl:value-of>
			</E2EDP15_TDNAME>
			<E2EDP15_SPRAS>
				<xsl:value-of select="'E'"></xsl:value-of>
			</E2EDP15_SPRAS>
			<E2EDP15_TXTLF>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXTLF>
			<E2EDP15_TXT01>
				<xsl:value-of select="'1260'"></xsl:value-of>
			</E2EDP15_TXT01>
			<E2EDP15_TXT02>
				<xsl:value-of select="concat(IMD_01/IMD_01_7008_01, IMD_01/IMD_01_7383)"/>
			</E2EDP15_TXT02>
			<E2EDP15_TXT03>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT03>
			<E2EDP15_TXT04>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT04>
			<E2EDP15_TXT05>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT05>
			<E2EDP15_TXT06>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT06>
			<E2EDP15_TXT07>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT07>
			<E2EDP15_TXT08>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT08>
			<E2EDP15_TXT09>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT09>
			<E2EDP15_TXT10>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT10>
			<E2EDP15_TXT11>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT11>
			<E2EDP15_TXT12>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT12>
			<E2EDP15_TXT13>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT13>
			<E2EDP15_TXT14>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT14>
			<E2EDP15_TSSPRAS_ISO>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TSSPRAS_ISO>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP15>
	</xsl:template>

	<!--Template for 'E2EDP15_YADT_1180' -->
	<xsl:template name="E2EDP15_YADT_1180">
		<E2EDP15>
			<E2EDP15_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SEGNAM>
			<E2EDP15_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_MANDT>
			<E2EDP15_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_DOCNUM>
			<E2EDP15_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_SEGNUM>
			<E2EDP15_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_PSGNUM>
			<E2EDP15_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_HLEVEL>
			<E2EDP15_TDNAME>
				<xsl:value-of select="'YADT'"></xsl:value-of>
			</E2EDP15_TDNAME>
			<E2EDP15_SPRAS>
				<xsl:value-of select="'E'"></xsl:value-of>
			</E2EDP15_SPRAS>
			<E2EDP15_TXTLF>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXTLF>
			<E2EDP15_TXT01>
				<xsl:value-of select="'1180'"></xsl:value-of>
			</E2EDP15_TXT01>
			<E2EDP15_TXT02>
				<xsl:value-of select="LIN_01/LIN_01_7140"/>
			</E2EDP15_TXT02>
			<E2EDP15_TXT03>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT03>
			<E2EDP15_TXT04>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT04>
			<E2EDP15_TXT05>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT05>
			<E2EDP15_TXT06>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT06>
			<E2EDP15_TXT07>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT07>
			<E2EDP15_TXT08>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT08>
			<E2EDP15_TXT09>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT09>
			<E2EDP15_TXT10>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT10>
			<E2EDP15_TXT11>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT11>
			<E2EDP15_TXT12>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT12>
			<E2EDP15_TXT13>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT13>
			<E2EDP15_TXT14>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TXT14>
			<E2EDP15_TSSPRAS_ISO>
				<!--Please insert the transformation instruction here-->
			</E2EDP15_TSSPRAS_ISO>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP15>
	</xsl:template>

	<!--Template for 'E2EDP16' -->
	<xsl:template name="E2EDP16">
		<E2EDP16>
			<E2EDP16_SEGNAM>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_SEGNAM>
			<E2EDP16_MANDT>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_MANDT>
			<E2EDP16_DOCNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_DOCNUM>
			<E2EDP16_SEGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_SEGNUM>
			<E2EDP16_PSGNUM>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_PSGNUM>
			<E2EDP16_HLEVEL>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_HLEVEL>
			<E2EDP16_ETTYP>
				<xsl:choose>
					<xsl:when test="SCC_01/SCC_01_4017 = '4'">
						<xsl:value-of select="'4'"></xsl:value-of>
					</xsl:when>
					<xsl:otherwise>
					</xsl:otherwise>
				</xsl:choose>
			</E2EDP16_ETTYP>
			<E2EDP16_PRGRS>
				<xsl:choose>
					<xsl:when test="SCC_01/SCC_01_2013 = 'Y'">
						<xsl:value-of select="'D'"></xsl:value-of>
					</xsl:when>
					<xsl:when test="SCC_01/SCC_01_2013 = 'W'">
						<xsl:value-of select="'W'"></xsl:value-of>
					</xsl:when>
					<xsl:when test="SCC_01/SCC_01_2013 = 'M'">
						<xsl:value-of select="'M'"></xsl:value-of>
					</xsl:when>
					<xsl:otherwise>
					</xsl:otherwise>
				</xsl:choose>
			</E2EDP16_PRGRS>
			<E2EDP16_EDATUV>
				<xsl:value-of select="DTM_08[DTM_08_2005 = '63']/DTM_08_2380"/>
			</E2EDP16_EDATUV>
			<E2EDP16_EZEIT>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_EZEIT>
			<E2EDP16_EDATUB>
				<xsl:value-of select="DTM_08[DTM_08_2005 = '63']/DTM_08_2380"/>
			</E2EDP16_EDATUB>
			<E2EDP16_ETVTF>
				<xsl:value-of select="'ZZ'"/>
			</E2EDP16_ETVTF>
			<E2EDP16_WMENG>
				<xsl:value-of select="format-number(QTY_01[QTY_01_6063 = '113']/QTY_01_6060, '#.####')"/>
			</E2EDP16_WMENG>
			<E2EDP16_FZABR>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_FZABR>
			<E2EDP16_BSTAS>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_BSTAS>
			<E2EDP16_WDATUV>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_WDATUV>
			<E2EDP16_WZEIT>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_WZEIT>
			<E2EDP16_WDATUB>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_WDATUB>
			<E2EDP16_LDATUV>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_LDATUV>
			<E2EDP16_LZEIT>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_LZEIT>
			<E2EDP16_LDATUB>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_LDATUB>
			<E2EDP16_BSTGRU>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_BSTGRU>
			<E2EDP16_TRANSPTP>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_TRANSPTP>
			<E2EDP16_TRANSPDESC>
				<!--Please insert the transformation instruction here-->
			</E2EDP16_TRANSPDESC>
			<rest>
				<!--Please insert the transformation instruction here-->
			</rest>
		</E2EDP16>
	</xsl:template>
</xsl:stylesheet>