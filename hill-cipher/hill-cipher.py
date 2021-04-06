import string
import numpy

# Use This For Encrypt
# 1. Encryption key : 3
# 2. Each row / matrix : 1 3 5, 1 2 3, 1 2 2
# 3. Encrypt
# 4. Enter Text

# Use this for decrypt
# 1. Encryption key : 3
# 2. Each row / matrix : 1 3 5, 1 2 3, 1 2 2
# 3. Decrypt
# 4. Enter text from encrypt result

# Referensi = http://python.algorithmexamples.com/web/ciphers/hill_cipher.html
 
 
def greatest_common_divisor(a: int, b: int) -> int:
    return b if a == 0 else greatest_common_divisor(b % a, a)
 
 
class HillCipher:
    key_string = string.ascii_uppercase + string.digits
    modulus = numpy.vectorize(lambda x: x % 36)
 
    to_int = numpy.vectorize(lambda x: round(x))
 
    def __init__(self, encrypt_key):
        self.encrypt_key = self.modulus(encrypt_key)  # mod36 calc's on the encrypt key
        self.check_determinant()  # validate the determinant of the encryption key
        self.decrypt_key = None
        self.break_key = encrypt_key.shape[0]
 
    def replace_letters(self, letter: str) -> int:
        return self.key_string.index(letter)
 
    def replace_digits(self, num: int) -> str:
        return self.key_string[round(num)]
 
    def check_determinant(self) -> None:
        det = round(numpy.linalg.det(self.encrypt_key))
 
        if det < 0:
            det = det % len(self.key_string)
 
        req_l = len(self.key_string)
        if greatest_common_divisor(det, len(self.key_string)) != 1:
            raise ValueError(
                f"determinant modular {req_l} of encryption key({det}) is not co prime w.r.t {req_l}.\nTry another key."
            )
 
    def process_text(self, text: str) -> str:
        chars = [char for char in text.upper() if char in self.key_string]
 
        last = chars[-1]
        while len(chars) % self.break_key != 0:
            chars.append(last)
 
        return "".join(chars)
 
    def encrypt(self, text: str) -> str:
        text = self.process_text(text.upper())
        encrypted = ""
 
        for i in range(0, len(text) - self.break_key + 1, self.break_key):
            batch = text[i : i + self.break_key]
            batch_vec = [self.replace_letters(char) for char in batch]
            batch_vec = numpy.array([batch_vec]).T
            batch_encrypted = self.modulus(self.encrypt_key.dot(batch_vec)).T.tolist()[
                0
            ]
            encrypted_batch = "".join(
                self.replace_digits(num) for num in batch_encrypted
            )
            encrypted += encrypted_batch
 
        return encrypted
 
    def make_decrypt_key(self):
        det = round(numpy.linalg.det(self.encrypt_key))
 
        if det < 0:
            det = det % len(self.key_string)
        det_inv = None
        for i in range(len(self.key_string)):
            if (det * i) % len(self.key_string) == 1:
                det_inv = i
                break
 
        inv_key = (
            det_inv
            * numpy.linalg.det(self.encrypt_key)
            * numpy.linalg.inv(self.encrypt_key)
        )
 
        return self.to_int(self.modulus(inv_key))
 
    def decrypt(self, text: str) -> str:
        self.decrypt_key = self.make_decrypt_key()
        text = self.process_text(text.upper())
        decrypted = ""
 
        for i in range(0, len(text) - self.break_key + 1, self.break_key):
            batch = text[i : i + self.break_key]
            batch_vec = [self.replace_letters(char) for char in batch]
            batch_vec = numpy.array([batch_vec]).T
            batch_decrypted = self.modulus(self.decrypt_key.dot(batch_vec)).T.tolist()[
                0
            ]
            decrypted_batch = "".join(
                self.replace_digits(num) for num in batch_decrypted
            )
            decrypted += decrypted_batch
 
        return decrypted
 
 
def main():
    N = int(input("Enter the order of the encryption key: "))
    hill_matrix = []
 
    print("Enter each row of the encryption key with space separated integers")
    for i in range(N):
        row = [int(x) for x in input().split()]
        hill_matrix.append(row)
 
    print(hill_matrix)
    hc = HillCipher(numpy.array(hill_matrix))
 
    print("Would you like to encrypt or decrypt some text? (1 or 2)")
    option = input("\n1. Encrypt\n2. Decrypt\n")
    if option == "1":
        text_e = input("What text would you like to encrypt?: ")
        print("Your encrypted text is:")
        print(hc.encrypt(text_e))
    elif option == "2":
        text_d = input("What text would you like to decrypt?: ")
        print("Your decrypted text is:")
        print(hc.decrypt(text_d))
 
 
if __name__ == "__main__":
    import doctest
 
    doctest.testmod()
 
    main()