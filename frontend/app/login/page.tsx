import Button from "@/components/button"

export default function login(): JSX.Element {
        return (
       <div>
            <h2 className="text-2xl flex font-bold place-content-center">Login</h2>
            <div className="m-4 flex flex-row h-screen justify-center">
               <input className="p-1 m-1 h-8 rounded" type="text" placeholder="Username" /> 
               <input className="p-1 m-1 h-8 rounded" type="text" placeholder="Password" /> 
               <Button link="/" text="Submit"/>
            </div>
       </div>
    )
}
