import ImgContainer from './imageContainer'
import Img from '@/types/Image'

type Props = {
        Imgs: Img[]
    }

export default function Gallery({Imgs}:Props): JSX.Element {
    
    if (!Imgs) {
            return (
            <h2 className="m-4 text-2xl font-bold"> No Images Found </h2>
            )
    }

    return (
        <>
            <section className="px-1 my-3 grid md:grid-cols-4 auto-rows">
                {Imgs.map(Img => (
                        <ImgContainer image={Img} key={Img.Filename} />
                    )
                )}
            </section>
        </>
    )
        
}
