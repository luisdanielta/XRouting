import ButtonBase from "./buttonBase";

interface SecondaryButtonProps {
  text: string;
  onClick?: () => void;
  disabled?: boolean;
}

const SecondaryButton: React.FC<SecondaryButtonProps> = ({
  text,
  onClick,
  disabled,
}) => {
  return (
    <ButtonBase
      text={text}
      onClick={onClick}
      disabled={disabled}
      className="shadow-gray-100 hover:shadow-gray-200 shadow outline-none focus:outline-none focus:ring-0 w-full bg-gray-200 transition-colors duration-200 ease-out hover:ease-in hover:bg-gray-300"
    />
  );
};

export default SecondaryButton;
