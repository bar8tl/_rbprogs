<?xml version="1.0" encoding="UTF-8"?>
<!--
#
# $Header:$
#
#	Mapping: IDOC YP3_DELVRY_RBNA -> EDIFACT DESDAV98A - VWAUDIMX
#   source description: oena_idoc_yp3_delvry_rbna_vw.fde	
#   target description: oena_edifact_desadv07a_vw.fde
#
# $Log:$
# 
# date		| user							| comment
# ..............................................................................................
# 20180716	| Anjana A   RBEI/BSF3		| initial mapping creation
#
-->
<!--Stylesheet was automatically generated on Mar 10, 2018 1:45:19 PM-->
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
	<xsl:import href="../tftab/OENA_SYS_IMP_UBK.xsl" />
	<xsl:import href="../tftab/OENA_CODETAB_OUTBOUND_UNB_CODES.xsl" />
	
	<!-- ############################################################### -->
	<!--  Parameters and Variables                                       -->
	<!-- ############################################################### -->

	<xsl:param name="ACTISEDI" select="''"/>
	<xsl:param name="PPATH" select="''"/>
	<xsl:param name="REF_NR" select="'000012345'"/>
	<xsl:param name="stepfile" select="''" />
	
	<!--  date, time from XSLT 2.0 functions-->
	<xsl:variable name="gv_docdate" select="format-date(current-date(),'[Y0001][M01][D01]')"/>
	<xsl:variable name="gv_edidate" select="format-date(current-date(),'[Y01][M01][D01]')" />
	<xsl:variable name="gv_editime" select="format-time(current-time(),'[H01][m01]')" />

	<!-- Required fields for Control segments -->
	<xsl:param name="UNB_S002_0004" select="''" saxon:assignable="yes"/>
	<!--   Syntax identifier-->
	<!-- <xsl:param name="UNB_S001_0001" select="'UNOC'"/> -->
		<!--  Sender identification						- LENGTH=35-->
	<xsl:param name="UNB_S002_0007" select="''" saxon:assignable="yes"/>
	<!--  Partner identification code qualifier		- LENGTH=4-->
	<xsl:param name="UNB_S002_0008"/>
	<!--  Address for reverse routing				- LENGTH=14-->
	<xsl:param name="UNB_S003_0010" select="''" saxon:assignable="yes"/>
	<!--  Recipient identification					- LENGTH=35-->
	<xsl:param name="UNB_S003_0007" select="''" saxon:assignable="yes"/>
	<!--  Partner identification code qualifier		- LENGTH=4-->
	<xsl:param name="UNB_S003_0014"/>
	<!-- Recipient's reference/password 			- LENGTH=14 -->
	<xsl:param name="UNB_S005_0022"/>
	<!-- Recipient's reference/password qualifier	- LENGTH=2 -->
	<xsl:param name="UNB_S005_0025"/>
	<!-- Application reference						- LENGTH=14 -->
	<xsl:param name="UNB_0026" />
	<!-- Processing priority code					- LENGTH=1 -->
	<xsl:param name="UNB_0029"/>
	<!-- Acknowledgement request					- LENGTH=1 -->
	<xsl:param name="UNB_0031"/>
	<!-- Communications agreement ID				- LENGTH=35 -->
	<xsl:param name="UNB_0032"/>
	<!--  Routing address							- LENGTH=14-->
	<xsl:param name="UNB_0035"/>
	
	<!-- saxon:assignable variables, need for counters -->
	<xsl:variable name="gv_lin_counter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_unh_counter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_message_reference_number" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_segment_counter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_st_counter" select="0" saxon:assignable="yes"/>
	<xsl:variable name="gv_PLT_COUNTER" select="0" saxon:assignable="yes"/>	
	<xsl:variable name="gv_PLT_FLAG" select="''" saxon:assignable="yes" />
	<xsl:variable name="gv_TEMP_CPS_CODE" select="''" saxon:assignable="yes" />
	<xsl:variable name="gv_TEMP_PAC_PCK" select="''" saxon:assignable="yes" />
	<xsl:variable name="gv_PALLET_CNT" select="0" saxon:assignable="yes" />
	<xsl:variable name="gv_ACCUM" select="0" saxon:assignable="yes" />
	<xsl:variable name="gv_TEMP_CPS_COUNT" select="0" saxon:assignable="yes" />
	
	
	<!--Other variables -->
	<xsl:variable name="gv_RCVLAD" select="''" saxon:assignable="yes"/>
	<xsl:variable name="gv_RCVPRN" select="/ACTIS/IDOC/EDI_DC/EDI_DC_RCVPRN" saxon:assignable="yes"/>
	<xsl:variable name="gv_EDI_DC_RCVLAD" select="tokenize(/ACTIS/IDOC/EDI_DC/EDI_DC_RCVLAD,';')" saxon:assignable="yes"/>
	<xsl:variable name="gv_ship_to" select="$gv_EDI_DC_RCVLAD[1]" saxon:assignable="yes"/>
	<xsl:variable name="gv_supplier_code" select="$gv_EDI_DC_RCVLAD[2]" saxon:assignable="yes"/>
	<xsl:variable name="gv_sold_to" select="$gv_EDI_DC_RCVLAD[3]" saxon:assignable="yes"/>
	<xsl:variable name="gv_SHIPPING_PLANT_ID" select="''" saxon:assignable="yes"/>
	<xsl:variable name="gv_SHIP_TO_CODE" select="''" saxon:assignable="yes"/>
	<xsl:variable name="gv_SOLD_TO_CODE" select="''" saxon:assignable="yes"/>
	<xsl:variable name="gv_SAP_SHIP_TO" select="''" saxon:assignable="yes"/>
	<xsl:variable name="gv_PICKUP_SCAC" select="''" saxon:assignable="yes"/>
	<xsl:variable name="gv_CARRIER_SCAC" select="''" saxon:assignable="yes"/>
	<xsl:variable name="gv_PAC_17" select="'FALSE'" saxon:assignable="yes" />
	<xsl:variable name="gv_TEMP_PAC" select="''" saxon:assignable="yes"/>
	<xsl:variable name="gv_TEMP_COUNTRY_OF_ORIGIN" select="'NO'" saxon:assignable="yes"/>
	<xsl:variable name="gv_TEMP_CITY1" select="''" saxon:assignable="yes"/>
	

	<!-- Mapping Name /-->
	<xsl:variable name="gv_mapping" select="'oena_idoc_yp3_delvry_rbna.oena_edifact_desadv07a_vwaudimx.xsl'" saxon:assignable="yes"/>

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
		
		<xsl:call-template name="M_DESADV_D07A" />

	</xsl:template>
	
	<!-- ################################ -->
	<!--  Create Structure                -->
	<!-- ################################ -->
	
	<!--Template for 'M_DESADV_D98A' -->
	<xsl:template name="M_DESADV_D07A">
	
			<!-- Get Customer Plant code -->
			<xsl:choose>
				<xsl:when test="(not(empty(/ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF = 'T1PC']/Z2Y04V_EDL24_ADI_VALUE)) and /ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF = 'T1PC']/Z2Y04V_EDL24_ADI_VALUE != '')">
					<saxon:assign name="gv_ship_to" select="/ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF = 'T1PC']/Z2Y04V_EDL24_ADI_VALUE"/>
				</xsl:when>
				<xsl:when test="(not(empty(/ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_KUWERK)) and /ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_KUWERK != '')">
					<saxon:assign name="gv_ship_to" select="/ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_KUWERK"/>
				</xsl:when>
				<xsl:otherwise>
					<xsl:message terminate="yes"><xsl:text>Mapping has been terminated because Customer_Plant_code is empty.</xsl:text></xsl:message>
				</xsl:otherwise>
			</xsl:choose>
			<!-- Get Supplier code -->
			<xsl:choose>
				<xsl:when test="(not(empty(/ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF = 'T1SC']/Z2Y04V_EDL24_ADI_VALUE)) and /ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF = 'T1SC']/Z2Y04V_EDL24_ADI_VALUE != '')">
					<saxon:assign name="gv_supplier_code" select="/ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF = 'T1SC']/Z2Y04V_EDL24_ADI_VALUE"/>
				</xsl:when>
				<xsl:when test="(not(empty(/ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_EIKTO_TRA)) and /ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_EIKTO_TRA != '')">
					<saxon:assign name="gv_supplier_code" select="/ACTIS/IDOC[1]/G01[1]/G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_EIKTO_TRA"/>
				</xsl:when>
				<xsl:otherwise>
					<xsl:message terminate="yes"><xsl:text>Mapping has been terminated because Supplier_code is empty.</xsl:text></xsl:message>
				</xsl:otherwise>
			</xsl:choose>
			<!-- Get Sold-to -->
			<xsl:choose>
				<xsl:when test="(not(empty(/ACTIS/IDOC[1]/G01[1]/G03/E2ADRM1[E2ADRM1_PARTNER_Q = 'AG']/E2ADRM1_PARTNER_ID)) and /ACTIS/IDOC[1]/G01[1]/G03/E2ADRM1[E2ADRM1_PARTNER_Q = 'AG']/E2ADRM1_PARTNER_ID != '')">
					<saxon:assign name="gv_sold_to" select="/ACTIS/IDOC[1]/G01[1]/G03/E2ADRM1[E2ADRM1_PARTNER_Q = 'AG']/E2ADRM1_PARTNER_ID"/>
				</xsl:when>
			</xsl:choose>
	
		<xsl:call-template name="OENA_SYS_IMP_UBK" />
		
		<xsl:message terminate="no"><xsl:text>RCVLAD: </xsl:text><xsl:value-of select="$gv_RCVLAD" />	</xsl:message>
		
		<saxon:assign name="UNB_S002_0004" select="$envelopeInfo/row[@Item eq $gv_RCVLAD][1]/@BOSCH_UNB_ID"/>
		<saxon:assign name="UNB_S002_0007" select="$envelopeInfo/row[@Item eq $gv_RCVLAD][1]/@BOSCH_UNB_QAULF"/>
		<saxon:assign name="UNB_S003_0010" select="$envelopeInfo/row[@Item eq $gv_RCVLAD][1]/@CUST_UNB_ID"/>
		<saxon:assign name="UNB_S003_0007" select="$envelopeInfo/row[@Item eq $gv_RCVLAD][1]/@CUST_UNB_QUALF"/>
		
		<xsl:choose>
			<xsl:when test="(not(empty($UNB_S002_0004)) and $UNB_S002_0004 != '')"/>
			<xsl:otherwise>
				<xsl:message terminate="yes">
					<xsl:text>ERROR: no value found in the OENA_SYS_IMP_UBK.xsl or OENA_CODETAB_OUTBOUND_UNB_CODES.xsl for </xsl:text>
					<xsl:text>Ship To : </xsl:text><xsl:value-of select="$gv_ship_to" />
					<xsl:text> , Supplier Code : </xsl:text><xsl:value-of select="$gv_supplier_code" />
					<xsl:text> , Sold To : </xsl:text><xsl:value-of select="$gv_sold_to" />
				</xsl:message>
			</xsl:otherwise>
		</xsl:choose>
		
		<M_DESADV_D07A>
			<xsl:call-template name="S_UNB" />
				<xsl:for-each select="/ACTIS/IDOC/G01/E2EDL20">
					<xsl:call-template name="SG0" />
				</xsl:for-each>
			<xsl:call-template name="S_UNZ" />
		</M_DESADV_D07A>
	</xsl:template>
	
	<!--Template for 'SG0' -->
	<xsl:template name="SG0">
	
		<!-- write Idoc number to log file -->
		<xsl:message terminate="no">
			<xsl:text>IDOC: </xsl:text>
			<xsl:value-of select="../../EDI_DC/EDI_DC_DOCNUM" />
		</xsl:message>
		
	
		<saxon:assign name="gv_st_counter" select="$gv_st_counter + 1" />
		<saxon:assign name="gv_segment_counter" select="0"/>
		<saxon:assign name="gv_lin_counter" select="0"/>
		<saxon:assign name="gv_PLT_FLAG" select="'N'"/>
		<saxon:assign name="gv_PALLET_CNT" select="0"/>
		<saxon:assign name="gv_ACCUM" select="0"/>
		<saxon:assign name="gv_TEMP_CPS_COUNT" select="0"/>
		
		<saxon:assign name="gv_RCVPRN" select="''"/>
		<saxon:assign name="gv_SHIPPING_PLANT_ID" select="''"/>
		<saxon:assign name="gv_SHIP_TO_CODE" select="''"/>
		<saxon:assign name="gv_SOLD_TO_CODE" select="''"/>
		<saxon:assign name="gv_SAP_SHIP_TO" select="''"/>
		<saxon:assign name="gv_TEMP_CPS_CODE" select="''"/>
		<saxon:assign name="gv_TEMP_PAC_PCK" select="''"/>
		<saxon:assign name="gv_RCVPRN" select="../../EDI_DC/EDI_DC_RCVPRN"/>
		<saxon:assign name="gv_SHIPPING_PLANT_ID" select="$gv_supplier_code"/>
		<saxon:assign name="gv_SHIP_TO_CODE" select="$gv_ship_to"/>
		<saxon:assign name="gv_SOLD_TO_CODE" select="$gv_sold_to"/>
		<saxon:assign name="gv_PICKUP_SCAC" select="''"/>
		<saxon:assign name="gv_CARRIER_SCAC" select="''"/>
		<saxon:assign name="gv_TEMP_PAC" select="''"/>
		<saxon:assign name="gv_TEMP_CITY1" select="''"/>
		
		<xsl:if test="../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UB']/E2ADRM1_PARTNER_ID != ''">
			<xsl:choose>
				<xsl:when test="contains(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UB']/E2ADRM1_PARTNER_ID,'-') ">
					<saxon:assign name="gv_CARRIER_SCAC" select="substring(substring-after(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UB']/E2ADRM1_PARTNER_ID,'-'),1,6)"/>
				</xsl:when>
				<xsl:otherwise>
					<saxon:assign name="gv_CARRIER_SCAC" select="substring(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UB']/E2ADRM1_PARTNER_ID,1,6)"/>
				</xsl:otherwise>
			</xsl:choose>
		</xsl:if>
		
		<xsl:if test="../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_PARTNER_ID != ''">
			<xsl:choose>
				<xsl:when test="contains(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_PARTNER_ID,'-') ">
					<saxon:assign name="gv_PICKUP_SCAC" select="substring(substring-after(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_PARTNER_ID,'-'),1,6)"/>
				</xsl:when>
				<xsl:otherwise>
					<saxon:assign name="gv_PICKUP_SCAC" select="substring(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_PARTNER_ID,1,6)"/>
				</xsl:otherwise>
			</xsl:choose>
		</xsl:if>
		
		<xsl:for-each select="../G03">
			<xsl:if test="./E2ADRM1/E2ADRM1_CITY1!=''">
				<saxon:assign name="gv_TEMP_CITY1" select="./E2ADRM1/E2ADRM1_CITY1"/>
			</xsl:if>
		</xsl:for-each>
		
		<xsl:if test="(not(empty(../G14[1]/G25/E2EDL35/E2EDL35_HERKL)) and ../G14[1]/G25/E2EDL35/E2EDL35_HERKL!= '')">	
			<saxon:assign name="gv_TEMP_COUNTRY_OF_ORIGIN" select="../G14[1]/G25/E2EDL35/E2EDL35_HERKL"/>
		</xsl:if>
		
		
		<SG0>
			<xsl:call-template name="S_UNH" />
			<xsl:call-template name="S_BGM" />
			<xsl:call-template name="S_DTM_545" />
			<xsl:call-template name="S_DTM_12" />
			<!-- <xsl:call-template name="S_DTM_549" /> -->
			<xsl:call-template name="S_DTM_550" />
			<xsl:call-template name="S_MEA" />
			<xsl:call-template name="S_MEA_69" />
			<xsl:call-template name="S_MEA_70" />
			<xsl:call-template name="SG1" />
			<xsl:call-template name="SG2" />
			<xsl:call-template name="SG6" />
			<xsl:call-template name="SG8" />
			<xsl:for-each select="../G14">
				<xsl:for-each select="Z2Y04V_DLITM">
					<xsl:if test="Z2Y04V_DLITM_QUAL = 'DUN' and Z2Y04V_DLITM_DATA4 = 'PLT'">
						<saxon:assign name="gv_PLT_FLAG" select="'Y'"/>
					</xsl:if>
				</xsl:for-each>
				
				<xsl:for-each select="./G15">
					<saxon:assign name="gv_PAC_17" select="'FALSE'"/>
					<xsl:call-template name="SG10_CPS" />
					<xsl:for-each select="./G16">
						<xsl:call-template name="SG10" />
					</xsl:for-each>
				</xsl:for-each>
			</xsl:for-each>				
			<xsl:call-template name="S_UNT" />
		</SG0>
	</xsl:template>
	
	<!--Template for 'SG1' -->
	<xsl:template name="SG1">
		<SG1>
			<xsl:call-template name="S_RFF_5" />
		</SG1>
	</xsl:template>

	<!--Template for 'SG2' -->
	<xsl:template name="SG2">
		<SG2>
			<xsl:call-template name="S_NAD_BY" />
		</SG2>
		<SG2>
			<xsl:call-template name="S_NAD_51" />
			<xsl:call-template name="SG3_70" />
		</SG2>
		<SG2>
			<xsl:call-template name="S_NAD_5" />
			<xsl:call-template name="SG3_8" />
		</SG2>
		<SG2>
			<xsl:call-template name="S_NAD_6" />
			<xsl:call-template name="S_LOC_14" />
		</SG2>
		<SG2>
			<xsl:call-template name="S_NAD" />
			<xsl:call-template name="SG3_69" />
			<!--<xsl:call-template name="SG3_2" />-->
		</SG2>
	</xsl:template>
	
	<!--Template for 'SG3_70' -->
	<xsl:template name="SG3_70">
		<SG3>
			<xsl:call-template name="S_RFF_70" />
		</SG3>
	</xsl:template>
	
	<!--Template for 'SG3_8' -->
	<xsl:template name="SG3_8">
		<SG3>
			<xsl:call-template name="S_RFF_8" />
		</SG3>
	</xsl:template>
	
	<!--Template for 'SG3_69' -->
	<xsl:template name="SG3_69">
		<SG3>
			<xsl:call-template name="S_RFF_69" />
		</SG3>
	</xsl:template>
	
	<!--Template for 'SG3_2' -->
	<xsl:template name="SG3_2">
		<SG3>
			<xsl:call-template name="S_RFF_2" />
		</SG3>
	</xsl:template>	
		
	
	<!--Template for 'SG6' -->
	<xsl:template name="SG6">
		<SG6>
			<xsl:call-template name="S_TDT" />
		</SG6>
	</xsl:template>
	
	<!--Template for 'SG8' -->
	<xsl:template name="SG8">
		<SG8>
			<xsl:call-template name="S_EQD" />
		</SG8>
	</xsl:template>
	
	<!--Template for 'SG10_CPS' -->
	<xsl:template name="SG10_CPS">
		<SG10>
			<xsl:call-template name="S_CPS_17" />
			<xsl:call-template name="SG11_GRP2" />
		</SG10>
	</xsl:template>
	
	<!--Template for 'SG10' -->
	<xsl:template name="SG10">
		<SG10>
			<xsl:for-each select="../Z2Y04V_AUXPKUNT">
				<xsl:if test="$gv_PAC_17 = 'FALSE'">
					<saxon:assign name="gv_PAC_17" select="'TRUE'"/>
					<xsl:call-template name="SG11_GRP1" />
				</xsl:if>	
			</xsl:for-each>	
			<xsl:call-template name="SG11_GRP3" />
			<xsl:for-each select="./G17">
				<xsl:call-template name="LIN_SG10" />
			</xsl:for-each>
		</SG10>
	</xsl:template>
	
	<!--Template for 'SG11_GRP1' -->
	<xsl:template name="SG11_GRP1">
		<SG11>
			<xsl:call-template name="S_PAC_17" />
		</SG11>
	</xsl:template>
	
	<!--Template for 'SG11_GRP2' -->
	<xsl:template name="SG11_GRP2">
		<SG11>
			<xsl:call-template name="S_PAC_21" />
			<xsl:call-template name="S_QTY_183" />
			<xsl:call-template name="SG13_GRP1" />
		</SG11>
	</xsl:template>

	<!--Template for 'SG13_GRP1' -->
	<xsl:template name="SG13_GRP1">
		<SG13>
			<xsl:call-template name="S_PCI_49" />
		</SG13>
	</xsl:template> 
	
	<!--Template for 'SG11_GRP3' -->
	<xsl:template name="SG11_GRP3">
		<SG11>
			<xsl:call-template name="SG13_GRP2" />
		</SG11>
	</xsl:template>

	<!--Template for 'SG13_GRP2' -->
	<xsl:template name="SG13_GRP2">
		<SG13>
			<xsl:call-template name="SG15_GRP1" />
			<!--<xsl:call-template name="SG15_GRP2" />-->
			<xsl:call-template name="SG15_GRP3" />
		</SG13>
	</xsl:template> 
	
	
	<!--Template for 'SG15_GRP1' -->
	<xsl:template name="SG15_GRP1">
		<SG15>
		<xsl:call-template name="S_GIN_178" />
		</SG15>
	</xsl:template>

	<!--Template for 'SG15_GRP2' -->
	<xsl:template name="SG15_GRP2">
		<SG15>
		<xsl:call-template name="S_GIN_159" />
		</SG15>
	</xsl:template>
	
	<!--Template for 'SG15_GRP3' -->
	<xsl:template name="SG15_GRP3">
		<SG15>
		<xsl:call-template name="S_GIN_159_2" />
		</SG15>
	</xsl:template>
	
	<!--Template for 'LIN_SG10' -->
	<xsl:template name="LIN_SG10">
		<LIN_SG10>
			<xsl:call-template name="LIN_S_CPS_21" />
			<xsl:call-template name="LIN_SG11_GRP1" />
			<xsl:call-template name="LIN_SG17" />
		</LIN_SG10>
	</xsl:template>
	
	<!--Template for 'LIN_SG11_GRP1' -->
	<xsl:template name="LIN_SG11_GRP1">
		<LIN_SG11>
			<xsl:call-template name="LIN_S_PAC_18" />
			<xsl:call-template name="LIN_S_QTY_184" />
			<xsl:call-template name="LIN_SG13_GRP1" />
		</LIN_SG11>
	</xsl:template>
	
	<!--Template for 'LIN_SG13_GRP1' -->
	<xsl:template name="LIN_SG13_GRP1">
		<LIN_SG13>
			<xsl:call-template name="LIN_S_PCI_53" />
			<xsl:for-each select="./Z2Y04V_SUBPACK">
				<xsl:call-template name="LIN_SG15_GRP1" />
			</xsl:for-each>
			<xsl:call-template name="LIN_SG15_GRP2" />
		</LIN_SG13>
	</xsl:template>
	
	<!--Template for 'LIN_SG15_GRP1' -->
	<xsl:template name="LIN_SG15_GRP1">
		<LIN_SG15>
		<xsl:call-template name="LIN_S_GIN_179" />
		</LIN_SG15>
	</xsl:template>

	<!--Template for 'LIN_SG15_GRP2' -->
	<xsl:template name="LIN_SG15_GRP2">
		<LIN_SG15>
		<xsl:call-template name="LIN_S_GIN_180" />
		</LIN_SG15>
	</xsl:template>
	
	<!--Template for 'LIN_SG17' -->
	<xsl:template name="LIN_SG17">
		<LIN_SG17>
			<xsl:call-template name="LIN_S_LIN_6" />
			<xsl:call-template name="LIN_S_PIA_6" />
			<xsl:call-template name="LIN_S_IMD_10" />
			<xsl:call-template name="LIN_S_QTY_187" />
			<xsl:call-template name="LIN_S_ALI_11" />
			<xsl:call-template name="LIN_SG18" />
			<xsl:call-template name="LIN_SG20" />
		</LIN_SG17>
	</xsl:template>
	
	<!--Template for 'LIN_SG18' -->
	<xsl:template name="LIN_SG18">
		<LIN_SG18>
			<xsl:call-template name="LIN_S_RFF_63" />
			<xsl:call-template name="LIN_S_DTM_651" />
		</LIN_SG18>
		<LIN_SG18>
			<xsl:call-template name="LIN_S_RFF_67" />
		</LIN_SG18>
	</xsl:template>
	
	<!--Template for 'LIN_SG20' -->
	<xsl:template name="LIN_SG20">
		<LIN_SG20>
			<!--<xsl:call-template name="LIN_S_LOC_24" />-->
		</LIN_SG20>
		<LIN_SG20>
			<xsl:call-template name="LIN_S_LOC_25" />
		</LIN_SG20>
	</xsl:template>
	
	<!-- ################################ -->
	<!--  Create Content                  -->
	<!-- ################################ -->

	<!-- ######################################## Interchange ######################################## -->

	<!--Template for 'S_UNB' -->
	<xsl:template name="S_UNB">
		<S_UNB>
			<idUNB></idUNB>
			<C_S001>
				<D_0001>
					<xsl:text>UNOC</xsl:text>
				</D_0001>
				<D_0002>
					<xsl:text>3</xsl:text>
				</D_0002>
			</C_S001>
			<C_S002>
				<D_0004>
					<xsl:value-of select="$UNB_S002_0004"/>
				</D_0004>
				<D_0007>
					<xsl:value-of select="$UNB_S002_0007"/>
				</D_0007>
				<D_0008>
					<xsl:value-of select="$UNB_S002_0008"/>
				</D_0008>
			</C_S002>
			<C_S003>
				<D_0010>
					<xsl:value-of select="$UNB_S003_0010"/>
				</D_0010>
				<D_0007>
					<xsl:value-of select="$UNB_S003_0007"/>
				</D_0007>
				<D_0014>
					<xsl:value-of select="$UNB_S003_0014"/>
				</D_0014>
			</C_S003>
			<C_S004>
				<D_0017>
					<xsl:value-of select="$gv_edidate"/>
				</D_0017>
				<D_0019>
					<xsl:value-of select="$gv_editime"/>
				</D_0019>
			</C_S004>
			<D_0020>
				<xsl:choose>
					<xsl:when test="$REF_NR != ''">
						<xsl:value-of select="format-number(856000000 + number($REF_NR),'000000000')"/>
					</xsl:when>
					<xsl:otherwise>
						<xsl:message terminate="no">
							<xsl:text>ERROR: Parameter $REF_NR not maintained in TSIM. Please fix!</xsl:text>
						</xsl:message>
					</xsl:otherwise>
				</xsl:choose>
			</D_0020>
			<C_S005>
				<D_0022></D_0022>
				<D_0025></D_0025>
			</C_S005>
			<D_0026>
				<xsl:value-of select="$UNB_0026"/>
			</D_0026>
			<D_0029>
				<xsl:value-of select="$UNB_0029"/>
			</D_0029>
			<D_0031>
				<xsl:value-of select="$UNB_0031"/>
			</D_0031>
			<D_0032>
				<xsl:value-of select="$UNB_0032"/>
			</D_0032>
			<D_0035>
				<xsl:value-of select="$UNB_0035"/>
			</D_0035>
		</S_UNB>
	</xsl:template>

	<!--Template for 'S_UNH' -->
	<xsl:template name="S_UNH">
		<xsl:call-template name="incrementSegmentCounter" />
		<xsl:call-template name="getNextUNH" />
		<saxon:assign name="gv_PLT_COUNTER" select="0"/>
		<saxon:assign name="gv_ACCUM" select="0"/>
		<saxon:assign name="gv_TEMP_CPS_COUNT" select="0"/>
		
		<S_UNH>
			<idUNH></idUNH>
			<D_0062>
				<xsl:value-of select="format-number(../../EDI_DC/EDI_DC_DOCNUM,'#')" />
			</D_0062>
			<C_S009>
				<D_0065>
					<xsl:text>DESADV</xsl:text>
				</D_0065>
				<D_0052>
					<xsl:text>D</xsl:text>
				</D_0052>
				<D_0054>
					<xsl:text>07A</xsl:text>
				</D_0054>
				<D_0051>
					<xsl:text>UN</xsl:text>
				</D_0051>
				<D_0057>
					<xsl:text>GAVF13</xsl:text>
				</D_0057>
			</C_S009>
			<D_0068></D_0068>
			<C_S010>
				<D_0070></D_0070>
				<D_0073></D_0073>
			</C_S010>
		</S_UNH>
	</xsl:template>
	
	<!--Template for 'S_UNT' -->
	<xsl:template name="S_UNT">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_UNT>
		<idUNT></idUNT>
			<D_0074>
				<xsl:value-of select="$gv_segment_counter"/>
			</D_0074>
			<D_0062>
				<xsl:value-of select="format-number(../../EDI_DC/EDI_DC_DOCNUM,'#')" />
			</D_0062>
		</S_UNT>
	</xsl:template>

	<!--Template for 'S_UNZ' -->
	<xsl:template name="S_UNZ">
		<S_UNZ>
			<idUNZ></idUNZ>
			<D_0036>
				<xsl:value-of select="$gv_unh_counter"/>
			</D_0036>
			<D_0020>
				<xsl:value-of select="format-number(856000000 + number($REF_NR),'000000000')"/>
			</D_0020>
		</S_UNZ>
	</xsl:template>
	
	<!--###############################################################-->
	<!-- Header                                                        -->
	<!--###############################################################-->

	<!--Template for 'S_BGM' -->
	<xsl:template name="S_BGM">
		<xsl:if test=" (not(empty(E2EDL20_VBELN)) and E2EDL20_VBELN != '') ">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_BGM>
				<idBGM></idBGM>
				<C_C002>
					<D_1001>
						<xsl:text>351</xsl:text>
					</D_1001>
					<D_1131></D_1131>
					<D_3055></D_3055>
					<D_1000>
						<xsl:text>LAB-ED</xsl:text>
					</D_1000>
				</C_C002>
				<C_C106>
					<D_1004>
						<xsl:if test=" (not(empty(E2EDL20_VBELN)) and E2EDL20_VBELN != '') ">
							<xsl:value-of select="substring(E2EDL20_VBELN,string-length(E2EDL20_VBELN)-7,8)"/>
						</xsl:if>
					</D_1004>
					<D_1056>
						<xsl:text>1</xsl:text>
					</D_1056>
					<D_1060></D_1060>
				</C_C106>
				<D_1225>
					<xsl:text>9</xsl:text>
				</D_1225>
				<D_4343></D_4343>
			</S_BGM>
		</xsl:if>
	</xsl:template>

	<!--Template for 'S_DTM_545' -->
	<xsl:template name="S_DTM_545">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_DTM>
			<idDTM></idDTM>
			<C_C507>
				<D_2005>
					<xsl:text>137</xsl:text>
				</D_2005>
				<D_2380>
					<xsl:value-of select="concat(../../EDI_DC/EDI_DC_CREDAT,substring(../../EDI_DC/EDI_DC_CRETIM,1,4))"/>
				</D_2380>
				<D_2379>
					<xsl:text>203</xsl:text>
				</D_2379>
			</C_C507>
		</S_DTM>
	</xsl:template>
	
	<!--Template for 'S_DTM_12' -->
	<xsl:template name="S_DTM_12">
		<xsl:if test="(not(empty(../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='DTM']/Z2Y04V_DLHDR_DATA1)) and ../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='DTM']/Z2Y04V_DLHDR_DATA1 != '')"> 
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_DTM>
				<idDTM></idDTM>
				<C_C507>
					<D_2005>
						<xsl:text>11</xsl:text>
					</D_2005>
					<D_2380>
						<xsl:value-of select="concat(../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='DTM']/Z2Y04V_DLHDR_DATA1,substring(../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='DTM']/Z2Y04V_DLHDR_DATA2,1,4))"/>
					</D_2380>
					<D_2379>
						<xsl:text>203</xsl:text>
					</D_2379>
				</C_C507>
			</S_DTM>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_DTM_549' -->
	<xsl:template name="S_DTM_549">
		<xsl:if test="../G14[1]/G19[1]/E2EDL52[E2EDL52_QUALF='DON']/E2EDL52_PRODN !=''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_DTM>
				<idDTM></idDTM>
				<C_C507>
					<D_2005>
						<xsl:text>191</xsl:text>
					</D_2005>
					<D_2380>
						<xsl:if test="../G14[1]/G19[1]/E2EDL52[E2EDL52_QUALF='DON']/E2EDL52_PRODN !=''">
							<xsl:value-of select="concat(substring(../G14[1]/G19[1]/E2EDL52[E2EDL52_QUALF='DON']/E2EDL52_PRODN,string-length(../G14[1]/G19[1]/E2EDL52[E2EDL52_QUALF='DON']/E2EDL52_PRODN)-7,8),'0000')"/>
						</xsl:if>
					</D_2380>
					<D_2379>
						<xsl:text>203</xsl:text>
					</D_2379>
				</C_C507>
			</S_DTM>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_DTM_550' -->
	<xsl:template name="S_DTM_550">
		<xsl:if test="(not(empty(../G04/E2EDT13[E2EDT13_QUALF='007']/E2EDT13_NTANF)) and ../G04/E2EDT13[E2EDT13_QUALF='007']/E2EDT13_NTANF != '')"> 
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_DTM>
				<idDTM></idDTM>
				<C_C507>
					<D_2005>
						<xsl:text>2</xsl:text>
					</D_2005>
					<D_2380>
						<xsl:value-of select="concat(../G04/E2EDT13[E2EDT13_QUALF='007']/E2EDT13_NTANF,substring(../G04/E2EDT13[E2EDT13_QUALF='007']/E2EDT13_NTANZ,1,4))"/>
					</D_2380>
					<D_2379>
						<xsl:text>203</xsl:text>
					</D_2379>
				</C_C507>
			</S_DTM>
		</xsl:if>
	</xsl:template>
	

	<!--Template for 'S_MEA' -->
	<xsl:template name="S_MEA">
		<xsl:if test="(not(empty(E2EDL20_BTGEW)) and E2EDL20_BTGEW != '')">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_MEA>
				<idMEA></idMEA>
				<D_6311>
					<xsl:text>AAX</xsl:text>
				</D_6311>
				<C_C502>
					<D_6313>
						<xsl:text>AAD</xsl:text>
					</D_6313>
					<D_6321></D_6321>
					<D_6155></D_6155>
					<D_6154></D_6154>
				</C_C502>
				<C_C174>
					<D_6411>
						<xsl:text>KGM</xsl:text>
					</D_6411>
					<D_6314>
						<xsl:if test="(not(empty(E2EDL20_BTGEW)) and E2EDL20_BTGEW != '')">
								<xsl:value-of select="format-number((E2EDL20_BTGEW),'#')"/>
						</xsl:if>					
					</D_6314>
					<D_6162></D_6162>
					<D_6152></D_6152>
					<D_6432></D_6432>
				</C_C174>
				<D_7383></D_7383>
			</S_MEA>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_MEA_69' -->
	<xsl:template name="S_MEA_69">
		<xsl:if test="(not(empty(E2EDL20_NTGEW)) and E2EDL20_NTGEW != '')">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_MEA>
				<idMEA></idMEA>
				<D_6311>
					<xsl:text>AAX</xsl:text>
				</D_6311>
				<C_C502>
					<D_6313>
						<xsl:text>AAL</xsl:text>
					</D_6313>
					<D_6321></D_6321>
					<D_6155></D_6155>
					<D_6154></D_6154>
				</C_C502>
				<C_C174>
					<D_6411>
						<xsl:text>KGM</xsl:text>
					</D_6411>
					<D_6314>
						<xsl:if test="(not(empty(E2EDL20_NTGEW)) and E2EDL20_NTGEW != '')">
							<xsl:value-of select="format-number((E2EDL20_NTGEW),'#')"/>
						</xsl:if>					
					</D_6314>
					<D_6162></D_6162>
					<D_6152></D_6152>
					<D_6432></D_6432>
				</C_C174>
				<D_7383></D_7383>
			</S_MEA>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_MEA_70' -->
	<xsl:template name="S_MEA_70">
		<xsl:if test="(not(empty(../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='PCK']/Z2Y04V_DLHDR_DATA2)) and ../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='PCK']/Z2Y04V_DLHDR_DATA2 != '')">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_MEA>
				<idMEA></idMEA>
				<D_6311>
					<xsl:text>AAE</xsl:text>
				</D_6311>
				<C_C502>
					<D_6313></D_6313>
					<D_6321></D_6321>
					<D_6155></D_6155>
					<D_6154></D_6154>
				</C_C502>
				<C_C174>
					<D_6411>
						<xsl:text>C62</xsl:text>
					</D_6411>
					<D_6314>
						<xsl:if test="(not(empty(../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='PCK']/Z2Y04V_DLHDR_DATA2)) and ../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='PCK']/Z2Y04V_DLHDR_DATA2 != '')">
								<xsl:value-of select="format-number(../Z2Y04V_DLHDR[Z2Y04V_DLHDR_QUAL='PCK']/Z2Y04V_DLHDR_DATA2,'#')"/>
						</xsl:if>					
					</D_6314>
					<D_6162></D_6162>
					<D_6152></D_6152>
					<D_6432></D_6432>
				</C_C174>
				<D_7383></D_7383>
			</S_MEA>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_RFF_5' -->
	<xsl:template name="S_RFF_5">
		<xsl:if test="(not(empty(E2EDL20_VBELN)) and E2EDL20_VBELN != '') and ../E2EDL22/E2EDL22_MANDT!=''">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_RFF>
			<idRFF></idRFF>
			<C_C506>
				<D_1153>
					<xsl:text>CRN</xsl:text>
				</D_1153>
				<D_1154>
					<xsl:choose>
						<xsl:when test="(not(empty(E2EDL20_VBELN)) and E2EDL20_VBELN != '')">
							<xsl:value-of select="substring(E2EDL20_VBELN,string-length(E2EDL20_VBELN)-7,8)"/>
						</xsl:when>
						<xsl:otherwise>
							<xsl:value-of select="../E2EDL22/E2EDL22_MANDT"/>
						</xsl:otherwise>
					</xsl:choose>
				</D_1154>
				<D_1156></D_1156>
				<D_4000></D_4000>
				<D_1060></D_1060>
			</C_C506>
		</S_RFF>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_NAD_BY' -->
	<xsl:template name="S_NAD_BY">
		<xsl:if test="../G03[E2ADRM1/E2ADRM1_PARTNER_Q='WE']/E2ADRE1/E2ADRE1_EXTEND_D !=''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_NAD>
				<idNAD></idNAD>
				<D_3035>
					<xsl:text>BY</xsl:text>
				</D_3035>
				<C_C082>
					<D_3039>
						<xsl:value-of select="../G03[E2ADRM1/E2ADRM1_PARTNER_Q='WE']/E2ADRE1/E2ADRE1_EXTEND_D"/>			
					</D_3039>
					<D_1131></D_1131>
					<D_3055>
						<xsl:text>92</xsl:text>
					</D_3055>
				</C_C082>
				<C_C058>
					<D_3124></D_3124>
				</C_C058>
				<C_C080>
					<D_3036></D_3036>
					<D_3045></D_3045>
				</C_C080>
				<C_C059>
					<D_3042></D_3042>
				</C_C059>
				<D_3164></D_3164>
				<C_C819>
					<D_3229></D_3229>
					<D_1131></D_1131>
					<D_3055></D_3055>
					<D_3228></D_3228>
				</C_C819>
				<D_3251></D_3251>
				<D_3207></D_3207>
			</S_NAD>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_NAD_51' -->
	<xsl:template name="S_NAD_51">
		<xsl:if test="(not(empty($gv_SHIPPING_PLANT_ID)) and $gv_SHIPPING_PLANT_ID != '')">	
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_NAD>
				<idNAD></idNAD>
				<D_3035>
					<xsl:text>SE</xsl:text>
				</D_3035>
				<C_C082>
					<D_3039>
						<xsl:value-of select="$gv_SHIPPING_PLANT_ID"/>				
					</D_3039>
					<D_1131></D_1131>
					<D_3055>
						<xsl:text>92</xsl:text>
					</D_3055>
				</C_C082>
				<C_C058>
					<D_3124></D_3124>
				</C_C058>
				<C_C080>
					<D_3036></D_3036>
					<D_3045></D_3045>
				</C_C080>
				<C_C059>
					<D_3042></D_3042>
				</C_C059>
				<D_3164></D_3164>
				<C_C819>
					<D_3229></D_3229>
					<D_1131></D_1131>
					<D_3055></D_3055>
					<D_3228></D_3228>
				</C_C819>
				<D_3251></D_3251>
				<D_3207></D_3207>
			</S_NAD>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_RFF_70' -->
	<xsl:template name="S_RFF_70">
		<xsl:if test="../G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_EIKTO_TRA!='' or 
					($gv_SHIPPING_PLANT_ID = '00X5022101'
					or $gv_SHIPPING_PLANT_ID = '00X7212801' 
					or $gv_SHIPPING_PLANT_ID = '00X7212803'
					or $gv_SHIPPING_PLANT_ID = '00X7212800'
					or $gv_SHIPPING_PLANT_ID = '00X7296400')">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_RFF>
			<idRFF></idRFF>
			<C_C506>
				<D_1153>
					<xsl:text>ANK</xsl:text>
				</D_1153>
				<D_1154>
					<xsl:choose>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X5022101'">
							<xsl:text>144845757</xsl:text>
						</xsl:when>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7212801'">
							<xsl:text>812047470</xsl:text>
						</xsl:when>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7212803'">
							<xsl:text>812857993</xsl:text>
						</xsl:when>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7212800'">
							<xsl:text>810248898</xsl:text>
						</xsl:when>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7296400'">
							<xsl:text>812813705</xsl:text>
						</xsl:when>
						<xsl:otherwise>
							<xsl:value-of select="../G14/Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF='DUNS']/Z2Y04V_EDL24_ADI_VALUE"/>
<!-- 							<xsl:value-of select="../G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_EIKTO_TRA"/> EDICUSTOMERREQUEST-2011 -->
						</xsl:otherwise>
					</xsl:choose>
				</D_1154>
				<D_1156></D_1156>
				<D_4000></D_4000>
				<D_1060></D_1060>
			</C_C506>
		</S_RFF>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_NAD_5' -->
	<xsl:template name="S_NAD_5">
		<xsl:if test="(not(empty($gv_SHIPPING_PLANT_ID)) and $gv_SHIPPING_PLANT_ID != '')">	
			<xsl:call-template name="incrementSegmentCounter"/>
				<S_NAD>
					<idNAD></idNAD>
					<D_3035>
						<xsl:text>SF</xsl:text>
					</D_3035>
					<C_C082>
						<D_3039>
							<xsl:value-of select="$gv_SHIPPING_PLANT_ID"/>
						</D_3039>
						<D_1131></D_1131>
						<D_3055>
							<xsl:text>92</xsl:text> 
						</D_3055>
					</C_C082>
					<C_C058>
						<D_3124></D_3124>
					</C_C058>
					<C_C080>
						<D_3036></D_3036>
						<D_3045></D_3045>
					</C_C080>
					<C_C059>
						<D_3042></D_3042>
					</C_C059>
					<D_3164></D_3164>
					<C_C819>
						<D_3229></D_3229>
						<D_1131></D_1131>
						<D_3055></D_3055>
						<D_3228></D_3228>
					</C_C819>
					<D_3251></D_3251>
					<D_3207></D_3207>
				</S_NAD>
			</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_RFF_8' -->
	<xsl:template name="S_RFF_8">
		<xsl:if test="../G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_EIKTO_TRA!='' or 
					($gv_SHIPPING_PLANT_ID = '00X5022101'
					or $gv_SHIPPING_PLANT_ID = '00X7212801' 
					or $gv_SHIPPING_PLANT_ID = '00X7212803'
					or $gv_SHIPPING_PLANT_ID = '00X7212800'
					or $gv_SHIPPING_PLANT_ID = '00X7296400')">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_RFF>
			<idRFF></idRFF>
			<C_C506>
				<D_1153>
					<xsl:text>ANK</xsl:text>
				</D_1153>
				<D_1154>
					<xsl:choose>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X5022101'">
							<xsl:text>144845757</xsl:text>
						</xsl:when>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7212801'">
							<xsl:text>812047470</xsl:text>
						</xsl:when>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7212803'">
							<xsl:text>812857993</xsl:text>
						</xsl:when>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7212800'">
							<xsl:text>810248898</xsl:text>
						</xsl:when>
						<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7296400'">
							<xsl:text>812813705</xsl:text>
						</xsl:when>
						<xsl:otherwise>
							<xsl:value-of select="../G14/Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF='DUNS']/Z2Y04V_EDL24_ADI_VALUE"/>
<!-- 							<xsl:value-of select="../G14[1]/Z2Y04V_EDL24_ASN/Z2Y04V_EDL24_ASN_EIKTO_TRA"/> EDICUSTOMERREQUEST-2011 -->
						</xsl:otherwise>
					</xsl:choose>
				</D_1154>
				<D_1156></D_1156>
				<D_4000></D_4000>
				<D_1060></D_1060>
			</C_C506>
		</S_RFF>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_NAD_6' -->
	<xsl:template name="S_NAD_6">
		<xsl:if test="(not(empty($gv_SHIP_TO_CODE)) and $gv_SHIP_TO_CODE != '')">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_NAD>
			<idNAD></idNAD>
			<D_3035>
				<xsl:if test="(not(empty($gv_SHIP_TO_CODE)) and $gv_SHIP_TO_CODE != '')">
					<xsl:text>ST</xsl:text>
				</xsl:if>
			</D_3035>
			<C_C082>
				<D_3039>
					<xsl:if test="(not(empty($gv_SHIP_TO_CODE)) and $gv_SHIP_TO_CODE != '')">
						<xsl:value-of select="$gv_SHIP_TO_CODE"/>
					</xsl:if>
				</D_3039>
				<D_1131></D_1131>
				<D_3055>
					<xsl:if test="(not(empty($gv_SHIP_TO_CODE)) and $gv_SHIP_TO_CODE != '')">
						<xsl:text>92</xsl:text>
					</xsl:if>
				</D_3055>
			</C_C082>
			<C_C058>
				<D_3124>
					<xsl:choose>
						<xsl:when test="(not(empty($gv_TEMP_CITY1)) and $gv_TEMP_CITY1 !='')">
							<xsl:value-of select="$gv_TEMP_CITY1"/>	
						</xsl:when>
					<xsl:otherwise>
						<xsl:value-of select="'CHATTANOOGA'"/>
					</xsl:otherwise>
					</xsl:choose>
				</D_3124>
				<D_3124></D_3124>
				<D_3124></D_3124>
				<D_3124></D_3124>
			</C_C058>
			<C_C080>
				<D_3036></D_3036>
				<D_3045></D_3045>
			</C_C080>
			<C_C059>
				<D_3042></D_3042>
			</C_C059>
			<D_3164></D_3164>
			<C_C819>
				<D_3229></D_3229>
				<D_1131></D_1131>
				<D_3055></D_3055>
				<D_3228></D_3228>
			</C_C819>
			<D_3251></D_3251>
			<D_3207></D_3207>
		</S_NAD>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_LOC_14' -->
	<xsl:template name="S_LOC_14">
		<xsl:if test="(not(empty(E2EDL20_ABLAD)) and E2EDL20_ABLAD != '')"> 
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_LOC>
			<idLOC></idLOC>
			<D_3227>
				<xsl:text>7</xsl:text>
			</D_3227>
			<C_C517>
				<D_3225>
					<xsl:if test="(not(empty(E2EDL20_ABLAD)) and E2EDL20_ABLAD != '')"> 
						<xsl:value-of select="E2EDL20_ABLAD"/>
					</xsl:if>
				</D_3225>
				<D_1131></D_1131>
				<D_3055>
					<xsl:text>92</xsl:text>
				</D_3055>
				<D_3224></D_3224>
			</C_C517>
			<C_C519>
				<D_3223></D_3223>
				<D_1131></D_1131>
				<D_3055></D_3055>
				<D_3222></D_3222>
			</C_C519>
			<C_C553>
				<D_3233></D_3233>
				<D_1131></D_1131>
				<D_3055></D_3055>
				<D_3232></D_3232>
			</C_C553>
			<D_5479></D_5479>
		</S_LOC>
		</xsl:if>
	</xsl:template>
	
	
	<!--Template for 'S_NAD' -->
	<xsl:template name="S_NAD">
		<xsl:if test="../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UF']/E2ADRM1_PARTNER_ID !='' or 
					($gv_SHIPPING_PLANT_ID = '00X5022101'
					or $gv_SHIPPING_PLANT_ID = '00X7212801' 
					or $gv_SHIPPING_PLANT_ID = '00X7212803'
					or $gv_SHIPPING_PLANT_ID = '00X7212800'
					or $gv_SHIPPING_PLANT_ID = '00X7296400')">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_NAD>
				<idNAD></idNAD>
				<D_3035>
					<xsl:text>FW</xsl:text>
				</D_3035>
				<C_C082>
					<D_3039>
						<xsl:choose>
							<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X5022101'">
								<xsl:text>ADP144845757</xsl:text>
							</xsl:when>
							<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7212801' 
											or $gv_SHIPPING_PLANT_ID = '00X7212803'
											or $gv_SHIPPING_PLANT_ID = '00X7212800'
											or $gv_SHIPPING_PLANT_ID = '00X7296400'">
								<xsl:text>DHL</xsl:text>
							</xsl:when>
							<xsl:otherwise>
								<xsl:if test="../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UF']/E2ADRM1_PARTNER_ID !=''">
									<xsl:value-of select="substring(substring-after(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UF']/E2ADRM1_PARTNER_ID, '-'),1,6)"/>
								</xsl:if>	
							</xsl:otherwise>
						</xsl:choose>
					</D_3039>
					<D_1131></D_1131>
					<D_3055>
						<xsl:text>92</xsl:text>
					</D_3055>
				</C_C082>
				<C_C058>
					<D_3124></D_3124>
				</C_C058>
				<C_C080>
					<D_3036></D_3036>
					<D_3045></D_3045>
				</C_C080>
				<C_C059>
					<D_3042></D_3042>
				</C_C059>
				<D_3164></D_3164>
				<C_C819>
					<D_3229></D_3229>
					<D_1131></D_1131>
					<D_3055></D_3055>
					<D_3228></D_3228>
				</C_C819>
				<D_3251></D_3251>
				<D_3207></D_3207>
			</S_NAD>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_RFF_69' -->
	<xsl:template name="S_RFF_69">
		<xsl:if test="($gv_SHIPPING_PLANT_ID = '00X5022101'
					or $gv_SHIPPING_PLANT_ID = '00X7212801' 
					or $gv_SHIPPING_PLANT_ID = '00X7212803'
					or $gv_SHIPPING_PLANT_ID = '00X7212800'
					or $gv_SHIPPING_PLANT_ID = '00X7296400')">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_RFF>
			<idRFF></idRFF>
			<C_C506>
				<D_1153>
					<xsl:text>ANK</xsl:text>
				</D_1153>
				<D_1154>
					<xsl:choose>
					<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X5022101'
									or $gv_SHIPPING_PLANT_ID = '00X7212801'
									or $gv_SHIPPING_PLANT_ID = '00X7212803'
									or $gv_SHIPPING_PLANT_ID = '00X7212800'">
						<xsl:text>812874071</xsl:text>
					</xsl:when>
					<xsl:when test=" $gv_SHIPPING_PLANT_ID = '00X7296400'">
						<xsl:text>812813705</xsl:text>
					</xsl:when>
					<xsl:otherwise>
					</xsl:otherwise>
					</xsl:choose>
				</D_1154>
				<D_1156></D_1156>
				<D_4000></D_4000>
				<D_1060></D_1060>
			</C_C506>
		</S_RFF>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_RFF_2' -->
	<xsl:template name="S_RFF_2">
		<xsl:if test=" (not(empty(E2EDL20_VBELN)) and E2EDL20_VBELN != '') ">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_RFF>
				<idRFF></idRFF>
				<C_C506>
					<D_1153>
						<xsl:text>ANK</xsl:text>
					</D_1153>
					<D_1154>
						<xsl:if test=" (not(empty(E2EDL20_VBELN)) and E2EDL20_VBELN != '') ">
							<xsl:value-of select="substring(E2EDL20_VBELN,string-length(E2EDL20_VBELN)-7,8)"/>
						</xsl:if>
					</D_1154>
					<D_1156></D_1156>
					<D_4000></D_4000>
					<D_1060></D_1060>
				</C_C506>
			</S_RFF>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_TDT' -->
	<xsl:template name="S_TDT">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_TDT>
			<idTDT></idTDT>
			<D_8051>
				<xsl:text>12</xsl:text>
			</D_8051>
			<D_8028></D_8028>
			<C_C220>
				<D_8067>
					<xsl:choose>
						<xsl:when test="normalize-space(substring(E2EDL20_TRATY,string-length(E2EDL20_TRATY)-1,2)) = 'TL'
										or normalize-space(substring(E2EDL20_TRATY,string-length(E2EDL20_TRATY)-1,2)) = 'TE'">
							<xsl:text>30</xsl:text>
						</xsl:when>
						<xsl:when test="normalize-space(substring(E2EDL20_TRATY,2,2)) = 'AF'">
							<xsl:text>40</xsl:text>
						</xsl:when>
						<xsl:otherwise>
							<xsl:if test="(not(empty(E2EDL20_TRATY)) and E2EDL20_TRATY != '')">
								<xsl:value-of select="normalize-space(substring(E2EDL20_TRATY,string-length(E2EDL20_TRATY)-1,2))"/>
							</xsl:if>
						</xsl:otherwise>
					</xsl:choose>
				</D_8067>
				<D_8066></D_8066>
			</C_C220>
			<C_C228>
				<D_8179></D_8179>
				<D_8178></D_8178>
			</C_C228>
			<C_C040>
				<D_3127>
					<xsl:if test="../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_PARTNER_ID !=''">
						<xsl:value-of select="substring(substring-after(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_PARTNER_ID, '-'),1,6)"/>
					</xsl:if>	
				</D_3127>
				<D_1131></D_1131>
				<D_3055>
					<xsl:text>92</xsl:text>
				</D_3055>
				<D_3126>
					<xsl:choose>
						<xsl:when test="../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_PARTNER_ID !=''">
							<xsl:value-of select="substring(substring-after(../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_PARTNER_ID, '-'),1,6)"/>
						</xsl:when>	
						<xsl:otherwise>
							<xsl:if test="../G03/E2ADRM1[E2ADRM1_PARTNER_Q='UU']/E2ADRM1_NAME1 = 'RYDR'">
								<xsl:text>RYDER</xsl:text>
							</xsl:if>
						</xsl:otherwise>
					</xsl:choose>
				</D_3126>
			</C_C040>
			<D_8101></D_8101>
			<C_C401>
				<D_8457>
					<xsl:text>ZZZ</xsl:text>
				</D_8457>
				<D_8459>
					<xsl:text>X</xsl:text>
				</D_8459>
				<D_7130>
					<xsl:text>X</xsl:text>
				</D_7130>
			</C_C401>
			<C_C222>
				<D_8213>
					<xsl:if test="(not(empty(../Z2Y04V_EDL20_ASN/Z2Y04V_EDL20_ASN_SIGNI)) and ../Z2Y04V_EDL20_ASN/Z2Y04V_EDL20_ASN_SIGNI != '')">
						<xsl:value-of select="substring(../Z2Y04V_EDL20_ASN/Z2Y04V_EDL20_ASN_SIGNI,1,9)"/>
					</xsl:if>
				</D_8213>
				<D_1131></D_1131>
				<D_3055>
					<!--<xsl:text>5</xsl:text>-->
				</D_3055>
				<D_8212></D_8212>
				<D_8453>
					<xsl:value-of select="$gv_TEMP_COUNTRY_OF_ORIGIN"/>
				</D_8453>
			</C_C222>
			<D_8281>
				<xsl:text>1</xsl:text>
			</D_8281>
		</S_TDT>
	</xsl:template>
	
	<!--Template for 'S_EQD' -->
	<xsl:template name="S_EQD">
		<xsl:if test="../Z2Y04V_EDL20_ASN/Z2Y04V_EDL20_ASN_SIGNI != '' or 
					E2EDL20_TRATY!=''">
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_EQD>
			<idEQD></idEQD>
			<D_8053>
				<xsl:choose>
					<xsl:when test=" normalize-space(substring(E2EDL20_TRATY,string-length(E2EDL20_TRATY)-1,2)) = 'TL' ">
						<xsl:text>TE</xsl:text>
					</xsl:when>
					<xsl:otherwise>
						<xsl:value-of select="normalize-space(substring(E2EDL20_TRATY,string-length(E2EDL20_TRATY)-1,2))"/>
					</xsl:otherwise>
				</xsl:choose>
			</D_8053>
			<C_C237>
				<D_8260>
					<xsl:if test="(not(empty(../Z2Y04V_EDL20_ASN/Z2Y04V_EDL20_ASN_SIGNI)) and ../Z2Y04V_EDL20_ASN/Z2Y04V_EDL20_ASN_SIGNI != '')">
						<xsl:value-of select="../Z2Y04V_EDL20_ASN/Z2Y04V_EDL20_ASN_SIGNI"/>
					</xsl:if>
				</D_8260>
				<D_1131></D_1131>
				<D_3055></D_3055>
				<D_3207></D_3207>
			</C_C237>
			<C_C224>
				<D_8155></D_8155>
				<D_1131></D_1131>
				<D_3055></D_3055>
				<D_8154></D_8154>
			</C_C224>
			<D_8077></D_8077>
			<D_8249></D_8249>
			<D_8169></D_8169>
		</S_EQD>
		</xsl:if>
	</xsl:template>
	
	<!--###############################################################-->
	<!-- Detail                                                        -->
	<!--###############################################################-->
	
	<!-- ######################################## Hierarchical Level-LIN ######################################## -->

	<!--Template for 'S_CPS_17' -->
	<xsl:template name="S_CPS_17">
		<saxon:assign name="gv_PLT_COUNTER" select="$gv_PLT_COUNTER + 1"/>
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_CPS>
			<idCPS></idCPS>
			<D_7164>
				<saxon:assign name="gv_TEMP_CPS_COUNT" select="$gv_PLT_COUNTER + $gv_ACCUM"/>
				<xsl:value-of select="$gv_TEMP_CPS_COUNT"/>
			</D_7164>
			<D_7166></D_7166>
			<D_7075>
				<xsl:if test="normalize-space(../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='BX'][1])"> 
					<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
				</xsl:if>
				<xsl:choose>
					<xsl:when test="$gv_PLT_FLAG = 'N'">
						<xsl:text>4</xsl:text>
					</xsl:when>
					<xsl:when test="$gv_TEMP_CPS_CODE = '1'">
						<xsl:text>3</xsl:text>
					</xsl:when>
					<xsl:otherwise>
						<xsl:text>4</xsl:text>
					</xsl:otherwise>
				</xsl:choose>
			</D_7075>
		</S_CPS>
	</xsl:template>

	<!--Template for 'S_PAC_17' -->
	<xsl:template name="S_PAC_17">	
		<xsl:call-template name="incrementSegmentCounter"/>
		<S_PAC>
			<idPAC></idPAC>
			<D_7224>
				<xsl:value-of select="format-number(Z2Y04V_AUXPKUNT_NO_OF_SHIPPING_UNITS,'#')"/>
			</D_7224>
			<C_C531>
				<D_7075></D_7075>
				<D_7233>
					<xsl:text>35</xsl:text>	
				</D_7233>
				<D_7073>
					<xsl:text>AAC</xsl:text>
				</D_7073>
			</C_C531>
			<C_C202>
				<D_7065>
					<xsl:value-of select="Z2Y04V_AUXPKUNT_PACKCODE"/>
				</D_7065>
				<D_1131></D_1131>
				<D_3055>
					<xsl:if test="Z2Y04V_AUXPKUNT_PACKCODE !=''">
						<xsl:text>92</xsl:text>
					</xsl:if>
				</D_3055>
				<D_7064></D_7064>
			</C_C202>
			<C_C402>
				<D_7077></D_7077>
				<D_7064></D_7064>
				<D_7143></D_7143>
			</C_C402>
			<C_C532>
				<D_8395></D_8395>
				<D_8393></D_8393>
			</C_C532>
		</S_PAC>
	</xsl:template>
	
	<!--Template for 'S_PAC_21' -->
	<xsl:template name="S_PAC_21">	
		<xsl:if test="$gv_PLT_FLAG != 'N'">
			<xsl:call-template name="incrementSegmentCounter"/>
			<xsl:if test="(not(empty(../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1)) and ../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1 !='')"> 
				<saxon:assign name="gv_TEMP_PAC_PCK" select="../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1"/>
			</xsl:if>
			<S_PAC>
				<idPAC></idPAC>
				<D_7224>
					<xsl:if test="normalize-space(../../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='BX'][1])"> 
						<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
					</xsl:if>
					<xsl:if test="$gv_PLT_FLAG != 'N'">
						<xsl:choose>
							<xsl:when test="$gv_TEMP_CPS_CODE = ''">
								<saxon:assign name="gv_PALLET_CNT" select="./Z2Y04V_PKUNT/Z2Y04V_PKUNT_NO_OF_SHIPPING_UNITS"/>
							</xsl:when>
							<xsl:otherwise>
								<saxon:assign name="gv_PALLET_CNT" select="1"/>
							</xsl:otherwise>
						</xsl:choose>
						<xsl:value-of select="$gv_PALLET_CNT"/>
					</xsl:if>
				</D_7224>
				<C_C531>
					<D_7075></D_7075>
					<D_7233>
						<xsl:text>35</xsl:text>	
					</D_7233>
					<D_7073>
						<xsl:text>AAC</xsl:text>
					</D_7073>
				</C_C531>
				<C_C202>
					<xsl:if test="$gv_TEMP_PAC_PCK !=''">
						<D_7065>
							<xsl:value-of select="$gv_TEMP_PAC_PCK"/>
						</D_7065>
						<D_1131></D_1131>
						<D_3055>
							<xsl:text>92</xsl:text>
						</D_3055>
						<D_7064></D_7064>
					</xsl:if>
				</C_C202>
				<C_C402>
					<D_7077></D_7077>
					<D_7064></D_7064>
					<D_7143></D_7143>
				</C_C402>
				<C_C532>
					<D_8395></D_8395>
					<D_8393></D_8393>
				</C_C532>
			</S_PAC>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_QTY_183' -->
	<xsl:template name="S_QTY_183">
		<xsl:if test="normalize-space(../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='BX'][1])"> 
			<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
		</xsl:if>
		<xsl:if test="$gv_TEMP_CPS_CODE = '' ">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_QTY>
				<idQTY></idQTY>
				<C_C186>
					<D_6063>
						<xsl:text>52</xsl:text>
					</D_6063>
					<D_6060>
						<xsl:if test="normalize-space(../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='BX'][1])"> 
							<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
						</xsl:if>
						<xsl:if test="$gv_PLT_FLAG != 'N'">
							<xsl:choose>
								<xsl:when test="$gv_TEMP_CPS_CODE = ''">
									<saxon:assign name="gv_PALLET_CNT" select="./Z2Y04V_PKUNT/Z2Y04V_PKUNT_NO_OF_SHIPPING_UNITS"/>
								</xsl:when>
								<xsl:otherwise>
									<saxon:assign name="gv_PALLET_CNT" select="1"/>
								</xsl:otherwise>
							</xsl:choose>
						</xsl:if>
						<xsl:choose>
							<xsl:when test="../E2EDL24/E2EDL24_LFIMG != '' and $gv_PALLET_CNT != 0">
								<xsl:value-of select="(../E2EDL24/E2EDL24_LFIMG div $gv_PALLET_CNT)"/>
							</xsl:when>
							<xsl:otherwise>
								<xsl:value-of select="format-number(../E2EDL24/E2EDL24_LFIMG,'#')"/>
							</xsl:otherwise>
						</xsl:choose>
					</D_6060>
					<D_6411>
						<xsl:text>PCE</xsl:text>
					</D_6411>
				</C_C186>
			</S_QTY>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_PCI_49' -->
	<xsl:template name="S_PCI_49">
		<xsl:if test="normalize-space(../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='BX'][1])"> 
			<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
		</xsl:if>
		<xsl:if test="(not(empty(../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1)) and ../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1 !='')"> 
			<saxon:assign name="gv_TEMP_PAC_PCK" select="../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1"/>
		</xsl:if>
		<xsl:if test="$gv_TEMP_PAC_PCK !=''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_PCI>
				<idPCI></idPCI>
				<D_4233>
					<xsl:text>17</xsl:text>
				</D_4233>
				<C_C210>
					<D_7102>
						<xsl:value-of select="substring(../../E2EDL20/E2EDL20_LIFEX,1,string-length(../../../../E2EDL20/E2EDL20_LIFEX)-9)"/>
					</D_7102>
				</C_C210>
				<D_8275></D_8275>
				<C_C827>
					<D_7511>
						<xsl:choose>
							<xsl:when test="$gv_TEMP_CPS_CODE = ''">
								<xsl:text>1J</xsl:text>
							</xsl:when>
							<xsl:otherwise>
								<xsl:text>6J</xsl:text>
							</xsl:otherwise>
						</xsl:choose>
					</D_7511>
					<D_1131></D_1131>
					<D_3055>
						<xsl:text>5</xsl:text>
					</D_3055>
				</C_C827>
			</S_PCI>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_GIN_178' -->
	<xsl:template name="S_GIN_178">
		<xsl:if test="(not(empty(./Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_SERNR)) and ./Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_SERNR !='')"> 
			<xsl:if test="(not(empty(../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1)) and ../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1 !='')"> 
				<saxon:assign name="gv_TEMP_PAC_PCK" select="../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='DUN'][Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1"/>
			</xsl:if>
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_GIN>
				<idGIN></idGIN>
				<D_7405>
					<xsl:if test="$gv_TEMP_PAC_PCK !=''">
						<xsl:text>ML</xsl:text>
					</xsl:if>
				</D_7405>
				<C_C208>
					<D_7402>
						<xsl:if test="$gv_PLT_FLAG != 'N'">
							<xsl:value-of select="./Z2Y04V_Z1EDL11[1]/Z2Y04V_Z1EDL11_SERNR"/>
						</xsl:if>
					</D_7402>
					<D_7402></D_7402>
				</C_C208>
			</S_GIN>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_GIN_159' -->
	<xsl:template name="S_GIN_159">
		<xsl:if test="(not(empty(../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1)) and ../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1 !='')"> 
			<saxon:assign name="gv_TEMP_PAC_PCK" select="../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='PLT']/Z2Y04V_DLITM_DATA1"/>
		</xsl:if>
		<xsl:if test="../Z2Y04V_DLITM/Z2Y04V_DLITM_DATA4='BX'"> 
			<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
		</xsl:if>
		<xsl:if test="$gv_TEMP_PAC_PCK !=''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_GIN>
				<idGIN></idGIN>
				<D_7405>
					<xsl:text>ML</xsl:text>
				</D_7405>
				<C_C208>
					<D_7402>
						<xsl:choose>
							<xsl:when test="./Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_SERNR !=''">
								<xsl:value-of select="./Z2Y04V_SUBPACK/Z2Y04V_Z1EDL11_SERNR"/>
							</xsl:when>
							<xsl:otherwise>
								<xsl:value-of select="./Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_SERNR"/>
							</xsl:otherwise>
						</xsl:choose>
					</D_7402>
					<D_7402></D_7402>
				</C_C208>
			</S_GIN>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'S_GIN_159_2' -->
	<xsl:template name="S_GIN_159_2">
		<xsl:if test="../../Z2Y04V_DLITM/Z2Y04V_DLITM_DATA4='BX'"> 
			<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
		</xsl:if>
		<xsl:if test="$gv_TEMP_CPS_CODE = ''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<S_GIN>
				<idGIN></idGIN>
				<D_7405>
					<xsl:text>BU</xsl:text>
				</D_7405>
				<C_C208>
					<D_7402>
						<xsl:choose>
							<xsl:when test="../../Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF='Z1'][last()]/Z2Y04V_EDL24_ADI_VALUE !='' ">
								<xsl:value-of select="../../Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF='Z1'][last()]/Z2Y04V_EDL24_ADI_VALUE"/>
							</xsl:when>
							<xsl:when test="string-length(../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='RAN']/Z2Y04V_DLITM_DATA1) &gt; 0 ">
								<xsl:value-of select="concat('00',substring(../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='RAN']/Z2Y04V_DLITM_DATA1,string-length(../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='RAN']/Z2Y04V_DLITM_DATA1)-7,8))"/>
							</xsl:when>
							<xsl:when test="../../G19[last()]/E2EDL52[E2EDL52_QUALF='DON']/E2EDL52_PRODN !=''">
								<xsl:value-of select="substring-before(../../G19[last()]/E2EDL52[E2EDL52_QUALF='DON']/E2EDL52_PRODN,';')"/>
							</xsl:when>
							<xsl:otherwise>
								<xsl:value-of select="./G16/Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_SERNR"/>
							</xsl:otherwise>
						</xsl:choose>
					</D_7402>
					<D_7402></D_7402>
				</C_C208>
			</S_GIN>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_CPS_21' -->
	<xsl:template name="LIN_S_CPS_21">
		<xsl:if test="../../../Z2Y04V_DLITM/Z2Y04V_DLITM_DATA4='BX'"> 
			<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
		</xsl:if>
		<saxon:assign name="gv_ACCUM" select="$gv_ACCUM + 1"/>
		<xsl:call-template name="incrementSegmentCounter"/>
		<LIN_S_CPS>
			<LIN_idCPS></LIN_idCPS>
			<LIN_D_7164>
				<xsl:value-of select="$gv_PLT_COUNTER + $gv_ACCUM"/>
			</LIN_D_7164>
			<LIN_D_7166>
				<xsl:if test="$gv_PLT_FLAG != 'N' and $gv_TEMP_CPS_CODE = '1'">
					<xsl:value-of select="$gv_TEMP_CPS_COUNT"/>
				</xsl:if>
			</LIN_D_7166>
			<LIN_D_7075>
				<xsl:if test="$gv_PLT_FLAG != 'N' and $gv_TEMP_CPS_CODE = '1'">
					<xsl:text>1</xsl:text>
				</xsl:if>
			</LIN_D_7075>
		</LIN_S_CPS>
	</xsl:template>
	
	<!--Template for 'LIN_S_PAC_18' -->
	<xsl:template name="LIN_S_PAC_18">	
		<xsl:if test="../../../Z2Y04V_DLITM/Z2Y04V_DLITM_DATA4='BX'"> 
			<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
		</xsl:if>
		<xsl:if test="../../../Z2Y04V_DLITM/Z2Y04V_DLITM_DATA4='BX'">
			<saxon:assign name="gv_TEMP_PAC" select="../../../Z2Y04V_DLITM[Z2Y04V_DLITM_DATA4='BX']/Z2Y04V_DLITM_DATA1"/>
		</xsl:if>
		
		<xsl:call-template name="incrementSegmentCounter"/>
		<LIN_S_PAC>
			<LIN_idPAC></LIN_idPAC>
			<LIN_D_7224>
				<xsl:value-of select="format-number(./Z2Y04V_SUBPACK[1]/Z2Y04V_SUBPACK_LFIMG,'#')"/>
			</LIN_D_7224>
			<LIN_C_C531>
				<LIN_D_7075></LIN_D_7075>
				<LIN_D_7233>
					<xsl:text>35</xsl:text>
				</LIN_D_7233>
				<LIN_D_7073>
					<xsl:text>AAC</xsl:text>
				</LIN_D_7073>
			</LIN_C_C531>
			<LIN_C_C202>
				<LIN_D_7065>
					<xsl:value-of select="$gv_TEMP_PAC"/>
				</LIN_D_7065>
				<LIN_D_1131></LIN_D_1131>
				<LIN_D_3055>
					<xsl:text>92</xsl:text>
				</LIN_D_3055>
				<LIN_D_7064></LIN_D_7064>
			</LIN_C_C202>
			<LIN_C_C402>
				<LIN_D_7077></LIN_D_7077>
				<LIN_D_7064></LIN_D_7064>
				<LIN_D_7143></LIN_D_7143>
			</LIN_C_C402>
			<LIN_C_C532>
				<LIN_D_8395></LIN_D_8395>
				<LIN_D_8393></LIN_D_8393>
			</LIN_C_C532>
		</LIN_S_PAC>
	</xsl:template>
	
	<!--Template for 'LIN_S_QTY_184' -->
	<xsl:template name="LIN_S_QTY_184">
		<xsl:if test="(not(empty(./Z2Y04V_SUBPACK[1]/Z2Y04V_SUBPACK_LFIMG)) and ./Z2Y04V_SUBPACK[1]/Z2Y04V_SUBPACK_LFIMG != '' )">
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_QTY>
				<LIN_idQTY></LIN_idQTY>
				<LIN_C_C186>
					<LIN_D_6063>
						<xsl:text>52</xsl:text>
					</LIN_D_6063>
					<LIN_D_6060>
						<xsl:if test="(not(empty(./Z2Y04V_SUBPACK[1]/Z2Y04V_SUBPACK_LFIMG)) and ./Z2Y04V_SUBPACK[1]/Z2Y04V_SUBPACK_LFIMG != '')">
							<xsl:value-of select="format-number(./Z2Y04V_SUBPACK[1]/Z2Y04V_SUBPACK_LFIMG,'#')"/>
						</xsl:if>
					</LIN_D_6060>
					<LIN_D_6411>
						<xsl:text>C62</xsl:text>
					</LIN_D_6411>
				</LIN_C_C186>
			</LIN_S_QTY>
		</xsl:if>
	</xsl:template>
	
		<!--Template for 'LIN_S_PCI_53' -->
	<xsl:template name="LIN_S_PCI_53">
		<xsl:if test="../../../Z2Y04V_DLITM/Z2Y04V_DLITM_DATA4='BX'"> 
			<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
		</xsl:if>
		<xsl:if test="$gv_TEMP_CPS_CODE != ''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_PCI>
				<LIN_idPCI></LIN_idPCI>
				<LIN_D_4233>
					<xsl:text>17</xsl:text>
				</LIN_D_4233>
				<LIN_C_C210>
					<LIN_D_7102>
						<xsl:if test="../../../../E2EDL20/E2EDL20_LIFEX != ''">
							<xsl:value-of select="substring(../../../../E2EDL20/E2EDL20_LIFEX,1,string-length(../../../../E2EDL20/E2EDL20_LIFEX)-9)"/>
						</xsl:if>
					</LIN_D_7102>
				</LIN_C_C210>
				<LIN_D_8275></LIN_D_8275>
				<LIN_C_C827>
						<LIN_D_7511>
							<xsl:text>1J</xsl:text>
						</LIN_D_7511>
						<LIN_D_1131></LIN_D_1131>
						<LIN_D_3055>
							<xsl:text>5</xsl:text>
						</LIN_D_3055>
				</LIN_C_C827>
			</LIN_S_PCI>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_GIN_179' -->
	<xsl:template name="LIN_S_GIN_179">
		<xsl:if test="Z2Y04V_SUBPACK_GERNR !=''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_GIN>
				<LIN_idGIN></LIN_idGIN>
				<LIN_D_7405>
					<xsl:text>ML</xsl:text>
				</LIN_D_7405>
				<LIN_C_C208>
					<LIN_D_7402>
						<xsl:value-of select="Z2Y04V_SUBPACK_GERNR"/>
					</LIN_D_7402>
					<LIN_D_7402></LIN_D_7402>
				</LIN_C_C208>
			</LIN_S_GIN>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_GIN_180' -->
	<xsl:template name="LIN_S_GIN_180">
		<xsl:if test="../../../Z2Y04V_DLITM/Z2Y04V_DLITM_DATA4='BX'"> 
			<saxon:assign name="gv_TEMP_CPS_CODE" select="'1'"/>
		</xsl:if>
		<xsl:if test="$gv_TEMP_CPS_CODE != '' and ../../../Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF='Z1'][last()]/Z2Y04V_EDL24_ADI_VALUE !=''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_GIN>
				<LIN_idGIN></LIN_idGIN>
				<LIN_D_7405>
					<xsl:text>BU</xsl:text>
				</LIN_D_7405>
				<LIN_C_C208>
					<LIN_D_7402>
						<xsl:choose>
						<xsl:when test=" ../../../Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF='Z1'][last()]/Z2Y04V_EDL24_ADI_VALUE !='' ">
							<xsl:value-of select="../../../Z2Y04V_EDL24_ADI[Z2Y04V_EDL24_ADI_QUALF='Z1'][last()]/Z2Y04V_EDL24_ADI_VALUE"/>
						</xsl:when>
						<!--<xsl:when test="string-length(../../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='RAN']/Z2Y04V_DLITM_DATA1) &gt; 0 ">
							<xsl:value-of select="concat('00',substring(../../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='RAN']/Z2Y04V_DLITM_DATA1,string-length(../../../Z2Y04V_DLITM[Z2Y04V_DLITM_QUAL='RAN']/Z2Y04V_DLITM_DATA1)-7,8))"/>
						</xsl:when>-->
						<!--<xsl:when test="../../../G19[last()]/E2EDL52[E2EDL52_QUALF='DON']/E2EDL52_PRODN !=''">
							<xsl:value-of select="substring-before(../../../G19[last()]/E2EDL52[E2EDL52_QUALF='DON']/E2EDL52_PRODN,';')"/>
						</xsl:when>-->
					</xsl:choose>
					</LIN_D_7402>
					<LIN_D_7402></LIN_D_7402>
				</LIN_C_C208>
			</LIN_S_GIN>
		</xsl:if>
	</xsl:template>
	

	<!-- ######################################## Hierarchical Level-Items ######################################## -->
	
	<!--Template for 'LIN_S_LIN_6' -->
	<xsl:template name="LIN_S_LIN_6">
		<xsl:if test="../../../E2EDL24/E2EDL24_KDMAT!=''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_LIN>
				<LIN_idLIN></LIN_idLIN>
				<LIN_D_1082>
				</LIN_D_1082>
				<LIN_D_1229></LIN_D_1229>
				<LIN_C_C212>
					<LIN_D_7140>
						<xsl:value-of select="concat('',../../../E2EDL24/E2EDL24_KDMAT)"/>
					</LIN_D_7140>
					<LIN_D_7143>
						<xsl:text>IN</xsl:text>
					</LIN_D_7143>
					<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055></LIN_D_3055>
				</LIN_C_C212>
				<LIN_C_C829>
					<LIN_D_5495></LIN_D_5495>
					<LIN_D_1082></LIN_D_1082>
				</LIN_C_C829>
				<D_1222></D_1222>
				<D_7083></D_7083>
			</LIN_S_LIN>
		</xsl:if>
	</xsl:template>

	<!--Template for 'LIN_S_PIA_6' -->
	<xsl:template name="LIN_S_PIA_6">
		<xsl:if test="(not(empty(../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR)) and ../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR!='') or ../Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_MANDT!=''"> 
		<xsl:call-template name="incrementSegmentCounter"/>
		<LIN_S_PIA>
			<LIN_idPIA></LIN_idPIA>
			<LIN_D_4347>
				<xsl:text>1</xsl:text>
			</LIN_D_4347>
			<LIN_C_C212>
				<LIN_D_7140>
					<xsl:choose>
						<xsl:when test="(not(empty(../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR)) and ../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR!='')"> 
							<xsl:value-of select="../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR"/>
						</xsl:when>
						<xsl:otherwise>
							<xsl:value-of select="../Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_MANDT"/>
						</xsl:otherwise>
					</xsl:choose>
				</LIN_D_7140>
				<LIN_D_7143>
					<xsl:text>SA</xsl:text>
				</LIN_D_7143>
				<LIN_D_1131></LIN_D_1131>
				<LIN_D_3055></LIN_D_3055>
			</LIN_C_C212>
			<LIN_C_C212>
				<LIN_D_7140></LIN_D_7140>
				<LIN_D_7143></LIN_D_7143>
				<LIN_D_1131></LIN_D_1131>
				<LIN_D_3055></LIN_D_3055>
			</LIN_C_C212>
		</LIN_S_PIA>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_IMD_10' -->
	<xsl:template name="LIN_S_IMD_10">
		<xsl:if test="(not(empty(../../../E2EDL24/E2EDL24_ARKTX)) and ../../../E2EDL24/E2EDL24_ARKTX != '') or ../Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_MANDT!=''">
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_IMD>
				<LIN_idIMD></LIN_idIMD>
				<LIN_D_7077>
					<xsl:text>C</xsl:text>
				</LIN_D_7077>
				<LIN_C_C272>
					<LIN_D_7081>
						<xsl:text>63</xsl:text>
					</LIN_D_7081>
					<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055></LIN_D_3055>
				</LIN_C_C272>
				<LIN_C_C273>
					<LIN_D_7009>
						<xsl:text>11</xsl:text>
					</LIN_D_7009>
					<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055>
						<xsl:text>272</xsl:text>
					</LIN_D_3055>
					<LIN_D_7008>
						<xsl:choose>
							<xsl:when test="(not(empty(../../../E2EDL24/E2EDL24_ARKTX)) and ../../../E2EDL24/E2EDL24_ARKTX != '')">
								<xsl:value-of select="substring(../../../E2EDL24/E2EDL24_ARKTX,1,35)"/>
							</xsl:when>
						<xsl:otherwise>
							<xsl:value-of select="../Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_MANDT"/>
						</xsl:otherwise>
					</xsl:choose>
					</LIN_D_7008>
					<LIN_D_3453></LIN_D_3453>
				</LIN_C_C273>
				<LIN_D_7383></LIN_D_7383>
			</LIN_S_IMD>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_QTY_187' -->
	<xsl:template name="LIN_S_QTY_187">
		<xsl:if test="(not(empty(../../../E2EDL24/E2EDL24_LFIMG)) and ../../../E2EDL24/E2EDL24_LFIMG != '')">
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_QTY>
				<LIN_idQTY></LIN_idQTY>
				<LIN_C_C186>
					<LIN_D_6063>
						<xsl:text>12</xsl:text>
					</LIN_D_6063>
					<LIN_D_6060>
						<xsl:value-of select="format-number(../../../E2EDL24/E2EDL24_LFIMG,'#')"/>
					</LIN_D_6060>
					<LIN_D_6411>
						<xsl:text>PCE</xsl:text>
					</LIN_D_6411>
				</LIN_C_C186>
			</LIN_S_QTY>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_ALI_11' -->
	<xsl:template name="LIN_S_ALI_11">
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_ALI>
				<LIN_idALI></LIN_idALI>
				<LIN_D_3239>
					<xsl:text>MX</xsl:text>
				</LIN_D_3239>
				<LIN_D_9213></LIN_D_9213>
				<LIN_D_4183></LIN_D_4183>
			</LIN_S_ALI>
	</xsl:template>
	
	<!--Template for 'LIN_S_RFF_63 -->
	<xsl:template name="LIN_S_RFF_63">
		<xsl:if test="(not(empty(../../../../E2EDL20/E2EDL20_VBELN)) and ../../../../E2EDL20/E2EDL20_VBELN!='')"> 
		<xsl:call-template name="incrementSegmentCounter"/>
		<LIN_S_RFF>
			<LIN_idRFF></LIN_idRFF>
			<LIN_C_C506>
				<LIN_D_1153>
					<xsl:text>AAU</xsl:text>
				</LIN_D_1153>
				<LIN_D_1154>
					<xsl:choose>
						<xsl:when test="(not(empty(../../../../E2EDL20/E2EDL20_VBELN)) and ../../../../E2EDL20/E2EDL20_VBELN!='')"> 
							<xsl:value-of select="substring(../../../../E2EDL20/E2EDL20_VBELN,string-length(../../../../E2EDL20/E2EDL20_VBELN)-7,8)"/>
						</xsl:when>
						<xsl:otherwise>
							<xsl:value-of select="../Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_MANDT"/>
						</xsl:otherwise>
					</xsl:choose>
				</LIN_D_1154>
				<LIN_D_1156>
					<xsl:text>1</xsl:text>
				</LIN_D_1156>
				<LIN_D_4000></LIN_D_4000>
				<LIN_D_1060></LIN_D_1060>
			</LIN_C_C506>
		</LIN_S_RFF>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_DTM_651' -->
	<xsl:template name="LIN_S_DTM_651">
		<xsl:call-template name="incrementSegmentCounter"/>
		<LIN_S_DTM>
			<LIN_idDTM></LIN_idDTM>
			<LIN_C_C507>
				<LIN_D_2005>
					<xsl:text>171</xsl:text>
				</LIN_D_2005>
				<LIN_D_2380>
					<xsl:value-of select="../../../../../EDI_DC/EDI_DC_CREDAT"/>
				</LIN_D_2380>
				<LIN_D_2379>
					<xsl:text>102</xsl:text>
				</LIN_D_2379>
			</LIN_C_C507>
		</LIN_S_DTM>
	</xsl:template>
	
	<!--Template for 'LIN_S_RFF_67 -->
	<xsl:template name="LIN_S_RFF_67">
		<xsl:if test="(not(empty(../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR)) and ../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR!='')"> 
		<xsl:call-template name="incrementSegmentCounter"/>
		<LIN_S_RFF>
			<LIN_idRFF></LIN_idRFF>
			<LIN_C_C506>
				<LIN_D_1153>
					<xsl:text>ON</xsl:text>
				</LIN_D_1153>
				<LIN_D_1154>
					<xsl:choose>
					<xsl:when test="(not(empty(../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR)) and ../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR!='')"> 
						<xsl:value-of select="../../../G26/E2EDL41[E2EDL41_QUALI='001']/E2EDL41_BSTNR"/>
					</xsl:when>
					<xsl:otherwise>
						<xsl:value-of select="../Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_MANDT"/>
					</xsl:otherwise>
					</xsl:choose>
				</LIN_D_1154>
				<LIN_D_1156></LIN_D_1156>
				<LIN_D_4000></LIN_D_4000>
				<LIN_D_1060></LIN_D_1060>
			</LIN_C_C506>
		</LIN_S_RFF>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_LOC_24' -->
	<xsl:template name="LIN_S_LOC_24">
		<xsl:if test="substring(../../../G27[1]/E2TXTP9/E2TXTP9_TDLINE,1,15) = 'UNLOADING POINT' or substring(../../../G27[1]/E2TXTP9/E2TXTP9_TDLINE,1,15) = 'unloading point' "> 
			<xsl:call-template name="incrementSegmentCounter"/>
			<LIN_S_LOC>
				<LIN_idLOC></LIN_idLOC>
				<LIN_D_3227>
					<xsl:text>11</xsl:text>
				</LIN_D_3227>
				<LIN_C_C517>
					<LIN_D_3225>
						<xsl:if test="substring(../../../G27[1]/E2TXTP9/E2TXTP9_TDLINE,1,15) = 'UNLOADING POINT' or substring(../../../G27[1]/E2TXTP9/E2TXTP9_TDLINE,1,15) = 'unloading point' "> 
							<xsl:value-of select="substring(substring-after(../../../G27[1]/E2TXTP9/E2TXTP9_TDLINE,':'),2,5)"/>
						</xsl:if>
					</LIN_D_3225>
					<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055>
						<xsl:if test="substring(../../../G27[1]/E2TXTP9/E2TXTP9_TDLINE,1,15) = 'UNLOADING POINT' or substring(../../../G27[1]/E2TXTP9/E2TXTP9_TDLINE,1,15) = 'unloading point' "> 
							<xsl:text>92</xsl:text>
						</xsl:if>
					</LIN_D_3055>
					<LIN_D_3224></LIN_D_3224>
				</LIN_C_C517>
				<LIN_C_C519>
					<LIN_D_3223></LIN_D_3223>
					<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055></LIN_D_3055>
					<LIN_D_3222></LIN_D_3222>
				</LIN_C_C519>
				<LIN_C_C553>
					<LIN_D_3233></LIN_D_3233>
					<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055></LIN_D_3055>
					<LIN_D_3232></LIN_D_3232>
				</LIN_C_C553>
				<LIN_D_5479></LIN_D_5479>
			</LIN_S_LOC>
		</xsl:if>
	</xsl:template>
	
	<!--Template for 'LIN_S_LOC_25' -->
	<xsl:template name="LIN_S_LOC_25">
	 <xsl:if test="(not(empty(../../../E2EDL24/E2EDL24_VBRST)) and ../../../E2EDL24/E2EDL24_VBRST != '') or ../Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_MANDT!=''  "> 
		<xsl:call-template name="incrementSegmentCounter"/>
		<LIN_S_LOC>
			<LIN_idLOC></LIN_idLOC>
			<LIN_D_3227>
				<xsl:text>11</xsl:text>
			</LIN_D_3227>
			<LIN_C_C517>
				<LIN_D_3225>
					<xsl:choose>
					<xsl:when test="(not(empty(../../../E2EDL24/E2EDL24_VBRST)) and ../../../E2EDL24/E2EDL24_VBRST != '')"> 
						<xsl:value-of select="../../../E2EDL24/E2EDL24_VBRST"/>
					</xsl:when>
					<xsl:otherwise>
						<xsl:value-of select="../Z2Y04V_Z1EDL11/Z2Y04V_Z1EDL11_MANDT"/>
					</xsl:otherwise>
					</xsl:choose>
				</LIN_D_3225>
				<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055>
						<xsl:text>92</xsl:text>
					</LIN_D_3055>
					<LIN_D_3224></LIN_D_3224>
				</LIN_C_C517>
				<LIN_C_C519>
					<LIN_D_3223></LIN_D_3223>
					<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055></LIN_D_3055>
					<LIN_D_3222></LIN_D_3222>
				</LIN_C_C519>
				<LIN_C_C553>
					<LIN_D_3233></LIN_D_3233>
					<LIN_D_1131></LIN_D_1131>
					<LIN_D_3055></LIN_D_3055>
					<LIN_D_3232></LIN_D_3232>
				</LIN_C_C553>
				<LIN_D_5479></LIN_D_5479>
			</LIN_S_LOC>
		</xsl:if>
	</xsl:template>
	
	
	
	<!--###############################################################-->
	<!-- END Mapping                                                   -->
	<!--###############################################################-->
	
</xsl:stylesheet>
