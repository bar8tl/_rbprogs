// replpdfd.java [2012-11-08 BAR8TL]
// Process user dialog box to specify XML invoice and PDF attachment paths into
// the PDF replacement (in base64 format) process
import java.awt.event.ActionListener;
import java.awt.event.ActionEvent;
import javax.swing.*;
import java.io.FileNotFoundException;
import java.io.IOException;
import javax.xml.parsers.ParserConfigurationException;
import javax.xml.xpath.XPathExpressionException;
import javax.xml.transform.TransformerConfigurationException;
import javax.xml.transform.TransformerException;
import org.xml.sax.SAXException;

public class replpdfd implements ActionListener {
  JLabel lblInvo1, lblInvo2, lblPdffl;
  JTextField txtInvo1, txtInvo2, txtPdffl;
  JButton btnInvo1, btnInvo2, btnPdffl, btnCancl, btnAcept;
  JFileChooser fchGenrl;

  public JPanel createContentPane ( ) {
    JPanel totalGUI = new JPanel( );
    totalGUI.setLocation(  0,  0);
    totalGUI.setSize(554,142);
    totalGUI.setLayout(null);

    lblInvo1 = new JLabel("Initial Invoice");
    lblInvo1.setLocation(  4, 10);
    lblInvo1.setSize( 74, 18);
    lblInvo1.setHorizontalAlignment(4);
    totalGUI.add(lblInvo1);

    lblPdffl = new JLabel("PDF File");
    lblPdffl.setLocation(  4, 30);
    lblPdffl.setSize( 74, 18);
    lblPdffl.setHorizontalAlignment(4);
    totalGUI.add(lblPdffl);

    lblInvo2 = new JLabel("Final Invoice");
    lblInvo2.setLocation(  4, 50);
    lblInvo2.setSize( 74, 18);
    lblInvo2.setHorizontalAlignment(4);
    totalGUI.add(lblInvo2);

    txtInvo1 = new JTextField(256);
    txtInvo1.setLocation( 86, 10);
    txtInvo1.setSize(360, 18);
    totalGUI.add(txtInvo1);

    txtPdffl = new JTextField(256);
    txtPdffl.setLocation( 86, 30);
    txtPdffl.setSize(360, 18);
    totalGUI.add(txtPdffl);

    txtInvo2 = new JTextField(256);
    txtInvo2.setLocation( 86, 50);
    txtInvo2.setSize(360, 18);
    totalGUI.add(txtInvo2);

    btnInvo1 = new JButton("Browse");
    btnInvo1.setLocation(450, 10);
    btnInvo1.setSize( 78, 18);
    btnInvo1.addActionListener(this);
    totalGUI.add(btnInvo1);

    btnPdffl = new JButton("Browse");
    btnPdffl.setLocation(450, 30);
    btnPdffl.setSize( 78, 18);
    btnPdffl.addActionListener(this);
    totalGUI.add(btnPdffl);

    btnInvo2 = new JButton("Browse");
    btnInvo2.setLocation(450, 50);
    btnInvo2.setSize( 78, 18);
    btnInvo2.addActionListener(this);
    totalGUI.add(btnInvo2);

    btnCancl = new JButton("Cancel");
    btnCancl.setLocation(198, 74);
    btnCancl.setSize(120, 22);
    btnCancl.addActionListener(this);
    totalGUI.add(btnCancl);

    btnAcept = new JButton("OK");
    btnAcept.setLocation(324, 74);
    btnAcept.setSize(120, 22);
    btnAcept.addActionListener(this);
    totalGUI.add(btnAcept);

    totalGUI.setOpaque(true);
    return totalGUI;
  }

  public void actionPerformed(ActionEvent e) {
    if (e.getSource() == btnInvo1) {
      fchGenrl = new JFileChooser();
      fchGenrl.setCurrentDirectory(new java.io.File("c:\\"));
      if (fchGenrl.showOpenDialog(null) == JFileChooser.APPROVE_OPTION)
        txtInvo1.setText(fchGenrl.getSelectedFile().toString());
    } else if (e.getSource() == btnPdffl) {
      fchGenrl = new JFileChooser();
      fchGenrl.setCurrentDirectory(new java.io.File("c:\\"));
      if (fchGenrl.showOpenDialog(null) == JFileChooser.APPROVE_OPTION)
        txtPdffl.setText(fchGenrl.getSelectedFile().toString());
    } else if (e.getSource() == btnInvo2) {
      fchGenrl = new JFileChooser();
      fchGenrl.setCurrentDirectory( new java.io.File( "c:\\" ) );
      if (fchGenrl.showSaveDialog(null) == JFileChooser.APPROVE_OPTION)
        txtInvo2.setText(fchGenrl.getSelectedFile().toString());
    } else if (e.getSource() == btnCancl) {
      System.exit(0);
    } else if (e.getSource() == btnAcept) {
      String iflnam = txtInvo1.getText();
      String pdfnam = txtPdffl.getText();
      String oflnam = txtInvo2.getText();
      replpdf repl = new replpdf();
      try {
        repl.EncodePDF(pdfnam).ReplPdfFile(iflnam, oflnam);
        JOptionPane.showMessageDialog(null, "PDF file has been replaced.",
          "Notification", JOptionPane.INFORMATION_MESSAGE);
      } catch ( FileNotFoundException ex ) {
        JOptionPane.showMessageDialog(null, "File not found.",
          "Error",JOptionPane.ERROR_MESSAGE);
      } catch ( IOException ex ) {
        JOptionPane.showMessageDialog(null, "IO Exception.",
          "Error", JOptionPane.ERROR_MESSAGE);
      } catch ( ParserConfigurationException ex ) {
        JOptionPane.showMessageDialog(null, "Parser Configuration Exception.",
          "Error", JOptionPane.ERROR_MESSAGE);
      } catch ( SAXException ex ) {
        JOptionPane.showMessageDialog(null, "SAX Exception.",
          "Error", JOptionPane.ERROR_MESSAGE);
      } catch ( XPathExpressionException ex ) {
        JOptionPane.showMessageDialog(null, "XPath Expression Exception.",
         "Error", JOptionPane.ERROR_MESSAGE);
      } catch ( TransformerConfigurationException ex ) {
       JOptionPane.showMessageDialog(null, "Transformer Config. Exception.",
         "Error", JOptionPane.ERROR_MESSAGE);
      } catch ( TransformerException ex ) {
       JOptionPane.showMessageDialog(null, "Transformer Exception.",
         "Error", JOptionPane.ERROR_MESSAGE);
      }
    }
  }

  private static void createAndShowGUI() {
    JFrame frame = new JFrame("Replace PDF File in VW eInvoices");
    replpdfd prog = new replpdfd();
    frame.setContentPane(prog.createContentPane());
    frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
    frame.setSize(554,142);
    frame.setLocation(200,200);
    frame.setVisible(true);
  }

  public static void main(String[] args) {
    SwingUtilities.invokeLater(new Runnable() {
      public void run() {
        createAndShowGUI();
      }
    });
  }
}
