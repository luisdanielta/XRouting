import ButtonBase from "./buttonBase";

interface PrimaryButtonProps {
  text: string;
  onClick?: () => void;
  disabled?: boolean;
}

const PrimaryButton: React.FC<PrimaryButtonProps> = ({
  text,
  onClick,
  disabled,
}) => {
  return (
    <ButtonBase
      text={text}
      onClick={onClick}
      disabled={disabled}
      className="bg-gradient-to-r from-blue-200 via-green-200 w-full to-green-300 text-gray-800 hover:text-gray-900 hover:bg-gradient-to-r hover:from-blue-300 hover:to-green-300 focus:ring-0 hover:ease-in shadow-gray-100 hover:shadow-gray-200"
    />
  );
};

export default PrimaryButton;
