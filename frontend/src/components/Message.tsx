export default function Message(props) {

    let getMessageClass = () => {
        switch (props.type) {
            case 'mine':
                return 'text-bg-primary'
            case 'bot':
                return 'text-bg-dark'
            default:
                return 'text-bg-secondary'
        }
    }

    return (
        <div class={getMessageClass() + ' p-2 mb-2'} style={'border-radius: 0.5em; font-size: 0.9em;'}>
            <div class="row">
                <div class="col" style={'font-weight: bold;'}>{props.name}</div>
                    <div class="col float-end text-end" style={'font-weight: bold; text-align: justify; text-justify: inter-word;'}>{props.time.toLocaleString()}</div>
            </div>
            <div class="row mt-2">
                <div class="col">
                    {props.msg}
                </div>
            </div>
        </div>
    )
}