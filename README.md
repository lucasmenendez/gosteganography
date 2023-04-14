# go-steganography
Simple implementation of LSB steganography algorithm


## Algorithm design

### Description
The algorithm uses the less important bits of every color of the input image pixels. 

Every image is composed by a matrix of pixels, and every pixel has, at least, three components, which are the value of the intensity of the three primary colors. This values can be any number between 0 and 255, and this range of values can be encoded in 8 bits:

    Pixel (12, 34) -> RGB(128, 92, 234)
        R: 128 -> 10000000
        G: 092 -> 01011100
        B: 234 -> 11101010

As we can see, in that values are some bits that are more *important* than others. For exampl, if we change the value of the first bit of the `92` number (`01011100` -> `11011100`), its decimal value changes to `220`. Otherwise, if we change the last bit (`01011100` -> `01011101`), its decimal value changes to `93`, which is different, but in terms of color intensity and image composition, it has not a perceptible change.

The algorithm uses these less important bits of every pixel, to hide the input message. This means that larger images can hide more information, and also that the size of the message is limited by the image size, trying to keep as much as posible the input image composition in the output one.

### Input components
 - **Image**: Where the message will be hide.
 - **Message**: Informtation that will be hide into the image.

### Basic steps
 1. Load the image from the source file and get its pixel matrix.
 2. Check if the message to be hide its smaller thant the number of pixels by 2. It is because pixel contains three bytes that contain the primary colors composition of this pixel. There are three primary colors and can be represented using any value between 0 and 255.
 3. Encode every pixel colors tuple (RGB) to its binary codification. Encode 
 also the message to binary.
 4. Iterate over the image pixels replacing the not modified less important bit
 of every color with a bit of the message until finish it. 
     - If the message length
 is greater than the number of pixels multiplied by 3 (the number of colors by pixel),
 iterate again over the image pixels replacing the sencond less important bit (instead the first one) by the next message bit until it will be fully included into the image.
