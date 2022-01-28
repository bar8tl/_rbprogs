// base64.java [2012-11-08 BAR8TL]
// Encode a bytes string into a base64 coded string
public class base64 {
  private final static byte[] _STANDARD_ALPHABET = {
    (byte)'A', (byte)'B', (byte)'C', (byte)'D', (byte)'E', (byte)'F',
    (byte)'G', (byte)'H', (byte)'I', (byte)'J', (byte)'K', (byte)'L',
    (byte)'M', (byte)'N', (byte)'O', (byte)'P', (byte)'Q', (byte)'R',
    (byte)'S', (byte)'T', (byte)'U', (byte)'V', (byte)'W', (byte)'X',
    (byte)'Y', (byte)'Z',
    (byte)'a', (byte)'b', (byte)'c', (byte)'d', (byte)'e', (byte)'f',
    (byte)'g', (byte)'h', (byte)'i', (byte)'j', (byte)'k', (byte)'l',
    (byte)'m', (byte)'n', (byte)'o', (byte)'p', (byte)'q', (byte)'r',
    (byte)'s', (byte)'t', (byte)'u', (byte)'v', (byte)'w', (byte)'x',
    (byte)'y', (byte)'z',
    (byte)'0', (byte)'1', (byte)'2', (byte)'3', (byte)'4', (byte)'5',
    (byte)'6', (byte)'7', (byte)'8', (byte)'9', (byte)'+', (byte)'/'
  };

  public static String encodeBytes(byte[] source) {
    String encoded = null;
    try {
      encoded = encodeBytes(source, 0, source.length);
    } catch (java.io.IOException ex) {
      assert false : ex.getMessage(); }
    assert encoded != null;
    return encoded;
  }

  public static String encodeBytes(byte[] source, int off, int len)
    throws java.io.IOException {
    byte[] encoded = encodeBytesToBytes(source, off, len);
    try {
      return new String(encoded, "US-ASCII");
    } catch (java.io.UnsupportedEncodingException uue) {
      return new String(encoded); }
  }

  public static byte[] encodeBytesToBytes(byte[] source, int off, int len)
    throws java.io.IOException {
    int encLen = (len/3)*4 + (len%3 > 0 ? 4 : 0);
    byte[] outBuff = new byte[encLen];

    int d = 0, e = 0, lineLength = 0, len2 = len-2;
    for (; d < len2; d+=3, e+=4) {
      encode3to4(source, d+off, 3, outBuff, e);
      lineLength += 4;
    }
    if (d < len) {
      encode3to4(source, d+off, len-d, outBuff, e);
      e += 4;
    }
    if (e <= outBuff.length-1) {
      byte[] finalOut = new byte[e];
      System.arraycopy(outBuff, 0, finalOut, 0, e);
      return finalOut;
    } else
      return outBuff;
  }

  private static byte[] encode3to4(byte[] source, int srcOffset,
    int numSigBytes, byte[] destination, int destOffset) {
    byte[] ALPHABET = _STANDARD_ALPHABET;
    int inBuff =
      (numSigBytes > 0 ? ((source[srcOffset    ] << 24) >>>  8) : 0) |
      (numSigBytes > 1 ? ((source[srcOffset + 1] << 24) >>> 16) : 0) |
      (numSigBytes > 2 ? ((source[srcOffset + 2] << 24) >>> 24) : 0);
    switch( numSigBytes ) {
      case 3:
        destination[destOffset    ] = ALPHABET[(inBuff >>> 18)       ];
        destination[destOffset + 1] = ALPHABET[(inBuff >>> 12) & 0x3f];
        destination[destOffset + 2] = ALPHABET[(inBuff >>>  6) & 0x3f];
        destination[destOffset + 3] = ALPHABET[(inBuff       ) & 0x3f];
        return destination;
      case 2:
        destination[destOffset    ] = ALPHABET[(inBuff >>> 18)       ];
        destination[destOffset + 1] = ALPHABET[(inBuff >>> 12) & 0x3f];
        destination[destOffset + 2] = ALPHABET[(inBuff >>>  6) & 0x3f];
        destination[destOffset + 3] = (byte)'=';
        return destination;
      case 1:
        destination[destOffset    ] = ALPHABET[(inBuff >>> 18)       ];
        destination[destOffset + 1] = ALPHABET[(inBuff >>> 12) & 0x3f];
        destination[destOffset + 2] = (byte)'=';
        destination[destOffset + 3] = (byte)'=';
        return destination;
      default:
        return destination;
    }
  }

  private base64(){}
}
