from PIL import Image
import pytesseract
import re

class CrosswordData:
    def __init__(self, left_word, right_word, length, marked_position):
        self.left_word = left_word
        self.right_word = right_word
        self.length = length
        self.marked_position = marked_position

    def __str__(self):
        return f"Left: {self.left_word}, Right: {self.right_word}, Length: {self.length}, Marked Position: {self.marked_position}"

def extract_text_from_image(image_path):
    return pytesseract.image_to_string(Image.open(image_path))

def parse_crossword_data(text):
    lines = [line.strip() for line in text.split("\n") if line.strip()]
    data = []

    for line in lines:
        match = re.match(r"^(\w+)\s+(\s+)\s+(\w+)$", line)
        if not match:
            continue

        left_word, space_between, right_word = match.groups()
        length = len(space_between)
        marked_position = space_between.find("^") + 1

        data.append(CrosswordData(left_word, right_word, length, marked_position))

    return data

if __name__ == "__main__":
    image_path = "test3.png"
    
    text = extract_text_from_image(image_path)
    print("Extracted Text:\n", text)

    extracted_data = parse_crossword_data(text)
    
    for data in extracted_data:
        print(data)



        