import Logo from "../components/Logo";

export default function Chats() {
    return (
        <div>
            <div class="modal" id="example-modal" tabindex="-1">
              <div class="modal-dialog">
                <div class="modal-content">
                  <div class="modal-header">
                    <h5 class="modal-title">Modal title</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                  </div>
                  <div class="modal-body">
                    <p>Modal body text goes here.</p>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                    <button type="button" class="btn btn-primary">Save changes</button>
                  </div>
                </div>
              </div>
            </div>
            <Logo />
            <div class="d-grid gap-2 mb-4" style={'text-align: center;'}>
                <p class="mt-4 display-6">Room List</p>
                <div class="list-group mt-4 mb-4">
                    <button type="button" class="list-group-item list-group-item-action list-group-item-lg">#1 Chat Room</button>
                </div>
                <button class="mt-4 btn btn-secondary btn-lg" data-bs-toggle="modal" data-bs-target="#example-modal">Add New</button>
                <button class="btn btn-white btn-lg">Logout</button>
            </div>
        </div>
    )
}