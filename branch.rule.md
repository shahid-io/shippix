# branch.rule.md

## Branching Strategy Guidelines for Shippix

A well-defined branching strategy will help your team organize work, enable effective code reviews, and streamline deployment. Follow these rules for consistent and successful collaboration on Shippix.

### 1. Main Branching Model

Adopt a simple and effective workflow such as [Git Flow](https://nvie.com/posts/a-successful-git-branching-model/) or [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow):

- **`main` (`master`) branch**  
  - Contains stable, production-ready code.
  - Only tested and reviewed code should be merged here.

- **`development` branch (optional)**  
  - Integrates features for the next release.
  - Useful if a staging branch/environment is needed.

- **Feature & Bugfix Branches**  
  - Each new feature or bugfix should have its own branch.
  - Branches are created from `develop` (or `main`, if not using `develop`).

    **Examples:**
    - `feature/shipment-creation`
    - `feature/auth-module`
    - `bugfix/login-error`

### 2. Branch Naming Conventions

- Use prefixes for clarity:
  - `feature/` – for new features
  - `bugfix/` – for bug fixes
- Branch names must be clear and descriptive.
- Optionally, include issue or ticket numbers for traceability.

    **Examples:**
    - `feature/shipment-123-add-tracking`
    - `bugfix/567-fix-auth-token`

### 3. Pull Requests & Code Reviews

- Always open a Pull Request (PR) from any feature or bugfix branch into `develop` or `main`.
- Use a PR template. Include:
    - Brief description
    - Tests performed
    - Screenshots (if applicable)
- PRs **must** be reviewed and approved by at least one team member before merging.

### 4. Continuous Integration & Deployment

- All PRs should trigger automated builds and tests.
- Use branch protection to prevent direct pushes to `main` (and `develop` if in use).
- Optionally, configure deployments for pre-release/staging from `develop`.

### 5. Advanced: Release & Hotfix Branches

- **Release branches:**  
  Prepare production releases without halting ongoing development.  
  Example: `release/1.2.0`
- **Hotfix branches:**  
  Address urgent production issues directly from `main`.  
  Example: `hotfix/production-crash-101`

## Summary

Following this branching strategy will help your team:

- Isolate development work
- Manage releases and emergency fixes
- Maintain a stable, production-ready codebase
- Support code reviews and parallel development

For Shippix’s backend microservices, using a workflow like Git Flow or GitHub Flow enhances team collaboration, quality, and confidence in deployment.

**Contributions and suggestions to this branching guideline are welcome.**

You can further edit this to match your internal documentation voice. Let me know if you need a version with more (or less) detail, or any additional sections (e.g., diagrams, examples)!