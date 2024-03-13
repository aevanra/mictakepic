import Socials from "@/components/socials"
import Button from "@/components/button"
import Gallery from "@/components/gallery"
import Img from "@/types/Image"


async function getHomeImages(): Promise<Img[]> {
        const res = await fetch('http://localhost:8082/listHomeImages');
        const data = await res.json();
        return data?.Images;
    }

export default async function Home(): Promise<JSX.Element> {
    const imageList = await getHomeImages()

    return (
        <div>
            <div className="text-end">
                <Button link="/login" text="User Login"/>
            </div>

            <Gallery Imgs={imageList} />

            <Socials/>

        </div>
    );
}
