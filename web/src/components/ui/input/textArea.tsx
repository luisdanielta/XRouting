import InputBase from "./inputBase";

interface TextAreaProps {
  placeholder?: string;
  className?: string;
  onChange?: (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => void | undefined;
  value?: string;
}

const TextArea: React.FC<TextAreaProps> = ({
  placeholder,
  className,
  onChange,
  value,
}) => {
  return (
    <InputBase
      type="textarea"
      placeholder={placeholder}
      className={`bg-white text-gray-700 ${className}`}
      onChange={onChange}
      value={value}
    />
  );
};

export default TextArea;
