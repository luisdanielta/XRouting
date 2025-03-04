import InputBase from "./inputBase";

interface InputProps {
  type?: "text" | "email" | "password";
  placeholder?: string;
  className?: string;
  onChange?: (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => void | undefined;
  value?: string;
}

const Input: React.FC<InputProps> = ({
  type = "text",
  placeholder,
  className,
  onChange,
  value,
}) => {
  return (
    <InputBase
      type={type}
      placeholder={placeholder}
      className={`bg-white text-gray-700 ${className}`}
      onChange={onChange}
      value={value}
    />
  );
};

export default Input;
