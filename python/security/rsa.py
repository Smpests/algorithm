import base64
from M2Crypto import RSA, m2, BIO


class RSA:
    """RSA加解密, PKCS8, 同java.security.spec.PKCS8EncodedKeySpec
    参考文档: https://blog.csdn.net/waitstory12/article/details/106377422
    """
    def __init__(self, private_key: str, public_key: str):
        """key为base64字符串, 之后可通过字符串格式加载不同的key"""
        self.private_key, self.public_key = self._load_key_pair(private_key, public_key)

    @staticmethod
    def _load_key_pair(private_keystring: str, public_keystring: str):
        private_key = RSA.load_key_string(f'-----BEGIN RSA PRIVATE KEY-----'
                                          f'\n{private_keystring}\n'
                                          f'-----END RSA PRIVATE KEY-----'.encode('utf-8'))
        public_key_bio = BIO.MemoryBuffer(f'-----BEGIN PUBLIC KEY-----'
                                          f'\n{public_keystring}\n'
                                          f'-----END PUBLIC KEY-----'.encode('utf-8'))
        public_key = RSA.load_pub_key_bio(public_key_bio)
        return private_key, public_key

    def private_encrypt(self, text):
        """私钥分段加密, 返回base64"""
        block_size = m2.rsa_size(self.private_key.rsa) - 11  # 加密时每个片段的最大长度（bytes）
        text_bytes = text.encode('utf-8')
        encrypt = b''
        for offset in range(0, len(text_bytes), block_size):
            segment = text_bytes[offset:offset + block_size]
            # 私钥加密，使用RSA.pkcs1_padding方式填充
            encrypt_seg = self.private_key.private_encrypt(segment, RSA.pkcs1_padding)
            encrypt += encrypt_seg
        return base64.b64encode(encrypt).decode('utf-8')

    def public_decrypt(self, base64_text) -> str:
        """公钥分段解密，返回明文字符串"""
        block_size = m2.rsa_size(self.public_key.rsa)
        text_bytes = base64.b64decode(base64_text)
        decrypt = b''
        for offset in range(0, len(text_bytes), block_size):
            segment = text_bytes[offset:offset + block_size]
            # 公钥解密，使用RSA.pkcs1_padding方式填充
            decrypt_seg = self.public_key.public_decrypt(segment, RSA.pkcs1_padding)
            decrypt += decrypt_seg
        return decrypt.decode('utf-8')
