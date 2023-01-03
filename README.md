# Observer perspective memory
## Program structure:
- **Entry point:** `EXECUTE.go`
- **Event control:** `event/keyPress.go`
- **Overall layout:** `obj/container.go`
    - **Object within each container:** `obj/canvas.go`

- **Materials control:** `material/files.go`
    - **Instructions:** `material/I/`
    - **Stimuli:** `material/P/`
- **Theme control:** `theme.go`

## Illustration:
The whole procedure is defined under `procedureController` func in `EXECUTE.go`. Adjusting it if there's any need. After making any change, type `go build -o EXECUTE.exe` in the terminal to rebuild the binary executable file.

Replacing the pictures in `material/P` or `material/I` to change the stimuli or instruction (don't need to rebuild).

Results will be exported to the `result` folder.

## Before executing:
Adjusting the parameters in the `.env` to what you want. Make sure to update the subject number(`SUBJECT_NUM`) and the assigned group(`GROUP`) before each task starts.

## Execution:
Double clicking `EXECUTE.exe` or typing `go run .` in the terminal to run this program.