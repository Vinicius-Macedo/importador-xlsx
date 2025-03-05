interface ContainerProps {
  children: React.ReactNode;
  className?: string;
}
export function Container(props: ContainerProps) {
  return (
    <div
      className={
        "w-full mx-auto lg:p-8 mb-12" +
        (props.className ? ` ${props.className}` : "")
      }
    >
      {props.children}
    </div>
  );
}
