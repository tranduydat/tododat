<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta content="width=device-width, initial-scale=1" name="viewport">
    <meta content="Just a todo app which made by Dat" name="description">
    <meta content="Dat Tran" name="author">
    <title>TodoDat</title>
    <link href="https://tododat.herokuapp.com" rel="canonical">

    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <style>
        .bd-placeholder-img {
            font-size: 1.125rem;
            text-anchor: middle;
            -webkit-user-select: none;
            -moz-user-select: none;
            user-select: none;
        }

        @media (min-width: 768px) {
            .bd-placeholder-img-lg {
                font-size: 3.5rem;
            }
        }
    </style>
</head>

<body class="bg-light">
<header class="py-4 text-center">
    <img alt="Logo" class="d-block mx-auto mb-3" src="/static/img/logo.svg" width="96">
    <h1 class="h2">TodoDat</h1>
</header>
<main>
    <section class="container">
    <form action="done" method="post">
        <div class="list-group">
            {{range $task := .doingTasks}}
            <label class="list-group-item">
                <input class="form-check-input me-1" onChange="this.form.submit()" name="taskid" type="checkbox" value="{{$task.Id}}">
                {{$task.Content}}
            </label>
            {{end}}
        </div>
        </form>
        <form action="add" method="post" class="mt-2 row g-2 justify-content-center">
                    <div class="col-8">
                        <input aria-describedby="button-addon2" aria-label="What task?" class="form-control" name="task-content"
                               placeholder="What task?" type="text">
                    </div>
                    <div class="col-2">
                        <input aria-describedby="button-addon2" aria-label="Due Date" class="form-control" name="task-due"
                               placeholder="Due Date" type="date">
                    </div>
                    <div class="col-1">
                        <select class="form-control form-select" id="inputState" name="task-priority">
                            <option selected>Priority</option>
                            <option value="1">1</option>
                            <option value="2">2</option>
                            <option value="3">3</option>
                            <option value="4">4</option>
                        </select>
                    </div>
                    <div class="col-1">
                        <button class="me-1 col-12 btn btn-outline-secondary" id="button-addon2" type="submit">Add</button>
                    </div>
                    </form>
        <div class="row justify-content-center mt-4">
            <div class="btn-group dropend col-2">
                <button class="btn btn-secondary" type="submit">
                    Menu
                </button>
                <button aria-expanded="false" class="btn btn-secondary dropdown-toggle dropdown-toggle-split"
                        data-bs-toggle="dropdown" type="button">
                    <span class="visually-hidden">More</span>
                </button>
                <ul class="dropdown-menu">
                    <li><button class="dropdown-item" formaction="edit" formmethod="post" type="submit">Edit</button></li>
                    <li><button class="dropdown-item" formaction="delete" formmethod="post" type="submit">Delete</button></li>
                    <li>
                        <hr class="dropdown-divider">
                    </li>
                    <li><a class="dropdown-item" href="#">My Account</a></li>
                    <li><a class="dropdown-item" href="#">Logout</a></li>
                </ul>
            </div>
        </div>
        </div>
    </form>
    </section>
</main>

<footer class="pt-5 text-muted text-center text-small">
    <p class="mb-1">&copy; 2021 Dat Tran</p>
</footer>
<script src="/static/js/bootstrap.bundle.min.js"></script>

<script src="form-validation.js"></script>
</body>
</html>