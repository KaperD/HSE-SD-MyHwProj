import tempfile
import subprocess

from dataclasses import dataclass
from pathlib import Path
from xmlrpc.client import Boolean


@dataclass
class Submission:
    solution_url: str


@dataclass
class Feedback:
    mark: int
    comment: str

    @classmethod
    def failed(cls, message: str):
        return Feedback(0, message)


@dataclass
class Howework:
    check_script_content: str

    def get_feedback(self, submission: Submission) -> Feedback:
        with tempfile.TemporaryDirectory() as runner_root_name:

            def exec(command: str, where: Path) -> subprocess.CompletedProcess:
                return subprocess.run(
                    command.split(),
                    cwd=where,
                    capture_output=True
                )

            def decode(b: bytes) -> str:
                return b.decode('utf-8')

            submission_source_dir_name = 'submission'
            submission_script_name = 'check.sh'
            runner_root = Path(runner_root_name)

            ret_git = exec(f'git clone --recursive {submission.solution_url} {submission_source_dir_name}',
                           runner_root)

            if ret_git.returncode != 0:
                return Feedback.failed(
                    f'''Failed to clone submission repo '{submission.solution_url}'
return code={ret_git.returncode}
stdout='{decode(ret_git.stdout)}'
stdout='{decode(ret_git.stderr)}'
'''
                )

            source_root = runner_root / submission_source_dir_name
            check_script_path = source_root / submission_script_name

            with check_script_path.open('w') as check_script_file:
                check_script_file.write(self.check_script_content)

            check = exec(f'sh {submission_script_name}', source_root)
            mark = 10 if check.returncode == 0 else 0
            comment = "" if check.stdout is None else decode(check.stdout)
            return Feedback(mark, comment)
