interface ButtonBaseProps {
  text: string;
  type?: "button" | "submit" | "reset";
  className?: string;
  onClick?: () => void;
  disabled?: boolean;
}

const ButtonBase: React.FC<ButtonBaseProps> = ({
  text,
  type = "submit",
  className = "",
  onClick,
  disabled = false,
}) => {
  return (
    <button
      type={type}
      className={`py-4 px-6 font-bold rounded focus:outline-none transition-colors duration-200 ease-out cursor-pointer shadow-md ${className}`}
      onClick={onClick}
      disabled={disabled}
    >
      {text}
    </button>
  );
};

export default ButtonBase;
