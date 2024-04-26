
interface Props {
        text: string;
    }

export default function Button(props: Props): JSX.Element {
    const defaultButtonClassName: string = "bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-2"
    return (
        <div>
            <button className={defaultButtonClassName}>
                {props.text}
            </button>
        </div>
    );
};
