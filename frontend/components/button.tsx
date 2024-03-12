import Link from "next/link";

interface Props {
        link: string;
        text: string;
    }

export default function Button(props: Props): JSX.Element {
    const defaultButtonClassName: string = "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-2"
    return (
        <div>
            <Link href={props.link}>
                <button className={defaultButtonClassName}>
                    {props.text}
                </button>
            </Link>
        </div>
    );
};
