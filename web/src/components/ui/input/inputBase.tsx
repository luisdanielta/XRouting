interface InputBaseProps {
  type?: "text" | "email" | "password" | "textarea";
  placeholder?: string;
  className?: string;
  onChange?: (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => void | undefined;
  value?: string;
}

const InputBase: React.FC<InputBaseProps> = ({
  type = "text",
  placeholder,
  className = "",
  onChange,
  value,
}) => {
  if (type === "textarea") {
    return (
      <textarea
        placeholder={placeholder}
        className={`rounded p-4 border-2 border-gray-200 hover:border-gray-300 shadow-gray-100 shadow outline-none focus:outline-none focus:ring-0 w-full bg-gray-50 transition-colors duration-200 ease-out hover:ease-in hover:bg-gray-100 ${className}`}
        onChange={onChange}
        value={value}
      />
    );
  }
  return (
    <input
      type={type}
      placeholder={placeholder}
      className={`rounded p-4 border-2 border-gray-200 hover:border-gray-300 shadow-gray-100 shadow outline-none focus:outline-none focus:ring-0 w-full bg-gray-50 transition-colors duration-200 ease-out hover:ease-in hover:bg-gray-100 ${className}`}
      onChange={onChange}
      value={value}
    />
  );
};

export default InputBase;
