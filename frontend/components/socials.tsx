import Image from "next/image";

export default function Social(): JSX.Element {
        return (
            <div className="m-4 flex flex-row place-content-center">
                <a href="https://www.instagram.com/mictakepic/" target="" className="p-1">
                    <Image src="/static/instagram_icon.png" alt="Instagram" className="social-icon" width={48} height={48}/>
                </a>
                <a href="https://www.youtube.com/@MicArmus" target="" className="p-1">
                    <Image src="/static/youtube_icon.png" alt="Youtube" className="social-icon" width={48} height={48}/>
                </a>
                <a href="https://www.tiktok.com/@miclarinet?lang=en" target="" className="p-1">
                    <Image src="/static/tiktok_icon.png" alt="Tiktok" className="social-icon" width={48} height={48}/>
                </a>
            </div>
    );
}

