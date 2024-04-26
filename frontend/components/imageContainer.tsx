import Image from "next/image"
import Img from "@/types/Image"

type Props = {
    image: Img; 
}

function buildImageSrc(image: Img): string {
    return "/Shares/" + image.DataShare + "/" + image.Filename
}

export default function ImageContainer({image}: Props) {

    const widthHeightRatio = image.Height / image.Width;
    const galleryHeight = Math.ceil(250*widthHeightRatio);
    const photoSpans = Math.ceil(galleryHeight / 10) + 1;

    return (

    <div className="justify-self-center" style= {{ gridRow: `span ${photoSpans}` }}>
        <div className="rounded-xl overflow-hidden m-1 place-content-center group">
            <Image
                src={buildImageSrc(image)}
                alt={image.Filename}
                width={image.Width}
                height={image.Height}
                sizes="250px"
                className="group-hover:opacity-75 "
            />
        </div>
    </div>

    );
}
