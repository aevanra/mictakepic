import Socials from "@/components/socials"
import Button from "@/components/button"
import Image from "next/image"

interface Image {
        Filename: string;
        Height: number;
        Width: number;
    } 

async function getHomeImages(): Promise<Image[]> {
        const res = await fetch('http://localhost:8082/listHomeImages');
        const data = await res.json();
        return data?.Images;
    }

function buildImageSrc(image: string): string {
        return "/Shares/micportfolio/" + image
    }

export default async function Home(): Promise<JSX.Element> {
    const imageList = await getHomeImages()

    return (
        <div>
            <div className="text-end">
                <Button link="/login" text="User Login"/>
            </div>

            <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                
                {imageList.map((image: Image) => (
                    <Image key={image.Filename} 
                        className="h-auto w-full relative rounded-lg" 
                        src={buildImageSrc(image.Filename)}
                        height={image.Height}
                        width={image.Width}
                        alt={buildImageSrc(image.Filename)}
                    />
                    ))}
            </div>

            <Socials/>

        </div>
    );
}
