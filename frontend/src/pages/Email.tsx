export default function Email(props) {
    return (
        <div>
            <div class="row text-center">
                <div class="col">
                    <p class="display-6">Email Sent</p>
                </div>
            </div>
            <div class="row mt-4">
                <div class="col">
                    <span>An email was sent to </span>
                    <span class="h6">{props.email}</span>
                    <span> please open it and confirm your sign up.</span>
                </div>
            </div>
            <div class="row mt-4 text-center">
                <div class="col">
                    <button class="btn btn-secondary">Sign In</button>
                </div>
            </div>
        </div>
    )
}