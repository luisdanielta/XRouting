
interface BtnProps {
    text: string
}

export default function Button({ text }: BtnProps) {

    return (
        <button>{text}</button>
    )
}