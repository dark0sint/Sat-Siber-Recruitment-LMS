package services

import (
    "os/exec"
    "strings"
)

func ProcessResume(resume string) string {
    // Call Python script
    cmd := exec.Command("python3", os.Getenv("PYTHON_SCRIPT_PATH")+"/resume_parser.py", resume)
    output, _ := cmd.Output()
    return strings.TrimSpace(string(output))
}
