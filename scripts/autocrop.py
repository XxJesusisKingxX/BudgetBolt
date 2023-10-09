from PIL import Image, ImageOps
import os
import argparse

# Input and output directories
parser = argparse.ArgumentParser(description='Auto-crop images in subfolders.')
parser.add_argument('inDir', help='Directory containing images')
parser.add_argument('outDir', help='Output directory for cropped images')
args = parser.parse_args()

# Create the output directory if it doesn't exist
os.makedirs(args.outDir, exist_ok=True)

# Loop through files in the input directory
for filename in os.listdir(args.inDir):
    if filename.endswith('.png') or filename.endswith('.jpg'):
        # Load the image
        input_path = os.path.join(args.inDir, filename)
        img = Image.open(input_path)
        
        # Remove the background (this assumes the background is white)
        # img = ImageOps.alpha_composite(Image.new('RGBA', img.size, (255, 255, 255, 255)), img)
        
        # Auto-crop the image
        img = img.crop(img.getbbox())
        
        # Save the cropped image to the output directory
        output_path = os.path.join(args.outDir, filename)
        img.save(output_path)
        print(f"Cropped and saved: {output_path}")
