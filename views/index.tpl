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
<header class="py-5 text-center">
    <img alt="Logo" class="d-block mx-auto mb-3" src="/static/img/logo.svg" width="96">
    <h1 class="h2">TodoDat</h1>
</header>
<main>
    <section class="container">
        <div class="row justify-content-md-center">
            <ul class="list-group">
            {{range $val := .tasks}}
                <li class="list-group-item">
                    <input aria-label="..." class="form-check-input me-1" type="checkbox" value="">
                    {{$val.Content}}
                </li>
              {{end}}
              <li class="list-group-item">
                <input aria-label="..." class="form-check-input me-1" type="checkbox" value="">

              </li>
            </ul>
            <button class="btn btn-primary btn mt-5 col-lg-2" type="submit">Update</button>
        </div>
    </section>
</main>

<footer class="my-5 pt-5 text-muted text-center text-small">
    <p class="mb-1">&copy; 2021 Dat Tran</p>
</footer>
<script src="/static/js/bootstrap.bundle.min.js"></script>

<script src="form-validation.js"></script>
</body>
</html>
