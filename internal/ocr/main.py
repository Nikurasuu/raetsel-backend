import cv2
import numpy as np
import pytesseract

image_path = "test.png"
image = cv2.imread(image_path)
height, width = image.shape[:2]

left_strip = image[:, :int(width * 0.2)]

left_text = pytesseract.image_to_string(left_strip, lang='deu')

lines = [line.strip() for line in left_text.split('\n') if line.strip()]
num_lines = len(lines)

print(num_lines)

results = []

for i in range(num_lines):
    y1 = int(i * height / num_lines)
    y2 = int((i + 1) * height / num_lines)
    line_img = image[y1:y2, :]

    text = pytesseract.image_to_string(line_img, lang='deu')
    words = text.strip().split()
    left_word = words[0] if len(words) > 0 else ""
    right_word = words[-1] if len(words) > 1 else ""

    h, w = line_img.shape[:2]
    crop = line_img[int(h*0.42):int(h*0.58), int(w*0.22):int(w*0.78)]
    gray = cv2.cvtColor(crop, cv2.COLOR_BGR2GRAY)
    blurred = cv2.GaussianBlur(gray, (5, 5), 0)

    num_cols = 111
    step = blurred.shape[1] // num_cols
    brightness = [np.mean(blurred[:, j*step:(j+1)*step]) for j in range(num_cols)]

    avg_brightness = np.mean(brightness)
    white_thresh = avg_brightness + 15

    best_count = -1
    best_binary = []

    for toleranz in range(0, 10):
        binary = [1 if b > (white_thresh - toleranz) else 0 for b in brightness]

        count = 0
        j = 0
        while j < len(binary):
            if binary[j] == 1:
                count += 1
                gap = 0
                while j < len(binary):
                    if binary[j] == 1:
                        gap = 0
                    else:
                        gap += 1
                    if gap > 0:
                        break
                    j += 1
            else:
                j += 1

        if count > best_count:
            best_count = count
            best_binary = binary.copy()

    results.append((left_word, right_word, best_count))

    debug_img = crop.copy()
    for j in range(num_cols):
        x = j * step
        color = (0, 255, 0) if best_binary[j] == 1 else (0, 0, 255)
        cv2.rectangle(debug_img, (x, 0), (x + step, crop.shape[0]), color, 1)
    cv2.imwrite(f"debug_line_{i+1}.png", debug_img)

for r in results:
    print(f"{r[0]}, {r[1]}, {r[2]}")