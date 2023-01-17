export default function RoomCreate() {
    return (
        <div class="text-center">
            <div class="row">
                <div class="col">
                    <p class="display-6">
                        New Room
                    </p>
                </div>
            </div>
            <div class="row mt-4">
                <div class="col d-grid gap-2">
                    <input class="form-control form-control-lg" type="text" placeholder="My Room" />
                    <button class="mt-4 btn btn-dark btn-lg">Confirm</button>
                    <button class="btn btn-white btn-lg">Cancel</button>
                </div>
            </div>
        </div>
        
    )
}