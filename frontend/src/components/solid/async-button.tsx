import { createSignal, type Component } from "solid-js"
import { Button, type buttonProps } from "../ui/button"

type AsyncButtonProps = {
  onClick: string | (() => Promise<any>);
  children: any;
} & Omit<buttonProps, 'onClick' | 'disabled'>;

const AsyncButton: Component<AsyncButtonProps> = (props) => {
  const [loading, setLoading] = createSignal(false)
  const fetchFunction = async (e: MouseEvent) => {
    e.preventDefault()
    setLoading(true)
    
    const handler = typeof props.onClick === 'string' 
      ? new Function('return ' + props.onClick)() 
      : props.onClick;
    console.log(handler, props.onClick)
    await handler()
    setLoading(false)
  }
  return (
    <Button disabled={loading()} {...props} onClick={fetchFunction}>
      {loading() ? (
        <div class="mr-2 h-4 w-4 animate-spin">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              class="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              stroke-width="4"
            ></circle>
            <path
              class="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            ></path>
          </svg>
        </div>
      ) : (
        props.children
      )}
    </Button>
  )
}

export default AsyncButton