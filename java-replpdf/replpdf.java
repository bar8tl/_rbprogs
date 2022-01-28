// replpdf.java [2012-11-08 BAR8TL]
// Main program to build and include a PDF attachmment into XML Mex invoices
import javax.xml.parsers.*;
import javax.xml.xpath.*;
import javax.xml.transform.*;
import javax.xml.transform.dom.DOMSource;
import javax.xml.transform.stream.StreamResult;
import org.xml.sax.SAXException;
import org.w3c.dom.*;
import java.io.*;

public class replpdf {
  String codb64 = null;

  public void ReplPdfFile(String ifil, String ofil)
    throws FileNotFoundException,IOException,ParserConfigurationException,
    SAXException,XPathExpressionException,TransformerConfigurationException,
    TransformerException {
    FileInputStream file = new FileInputStream(new File(ifil));
    DocumentBuilderFactory fact = DocumentBuilderFactory.newInstance();
    DocumentBuilder bldr = fact.newDocumentBuilder();
    Document docm = bldr.parse(file);
    XPath xpath = XPathFactory.newInstance().newXPath();
    Node node = (Node) xpath.compile("/Comprobante/Addenda/Factura/Archivo").
      evaluate(docm, XPathConstants.NODE);
    if (node != null) {
      ((Element) node).setAttribute("datos", new String(codb64));
      Transformer tran = TransformerFactory.newInstance().newTransformer();
      StreamResult output = new StreamResult(new File(ofil));
      DOMSource input  = new DOMSource(docm);
      tran.transform(input, output);
    }
  }

  public replpdf EncodePDF(String ifname)
    throws FileNotFoundException,IOException {
    File ifile = new File(ifname);
    byte[] result = new byte[(int)ifile.length()];
    InputStream istrm = new BufferedInputStream(new FileInputStream(ifile));
    istrm.read(result, 0, (int)ifile.length());
    istrm.close();
    codb64 = base64.encodeBytes(result);
    return this;
  }
}
