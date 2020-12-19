<?php

function getAllVictims(){
    # validate input for sql injection
    # end points should response only local requests except few of them
    $query = "SELECT * FROM ...";

    $infos = array(
        array("id"=>"1", "ip"=>"192.168.1.1", "username"=>"okan", "infection_date"=>"19/12/2020 21:33", "first_touch"=> "20/12/2020 09=>33", "country"=>"Turkey"),
        array("id"=>"3", "ip"=>"12.32.12.1", "username"=>"ogulcan", "infection_date"=>"11/11/2020 21:33", "first_touch"=> "11/11/2020 21=>43", "country"=>"Turkey"),
        array("id"=>"4", "ip"=>"34.43.15.12", "username"=>"idil", "infection_date"=>"19/12/2020 11:33", "first_touch"=> "20/12/2020 19=>33", "country"=>"Germany")
    );

    return json_encode($infos);
}

function createVictimTable(){
    $victims = json_decode(getAllVictims());
    //print_r($victims);
    $table = '<table class="table table-bordered">';
    $table.= '<thead>
                    <tr>
                    <th scope="col">ID</th>
                    <th scope="col">IP</th>
                    <th scope="col">Username</th>
                    <th scope="col">Infection Date</th>
                    <th scope="col">First Touch Date</th>
                    <th scope="col">Country</th>
                    </tr>
                </thead>
                <tbody>';
    
    foreach($victims as $vic){
        $table.='<tr>';
        foreach($vic as $ky => $vl){
            #echo $ky."=>".$vl."<br>";
            $table.='<td>'.$vl.'</td>';
        }
        $table.='</tr>';
    }

    $table.='</tbody>
            </table>';
    
    return $table;
}

?>


<!-- HTML -->

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>APT_Al</title>
    <!-- CSS -->
    <link rel="stylesheet" href="/css/style.css">
    <link rel="stylesheet" href="/css/bootstrap.min.css">

</head>
    <body>

        <div class="topnav">
            <nav class="navbar navbar-expand-lg navbar-dark" style="background-color: #070812">
                <a class="navbar-brand" href="/Status"><img src="/image/logo1.png" style="height: 60px; width: auto"></a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>

                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav mr-auto">
                        <li class="nav-item">
                            <a class="nav-link" href="/Status/index.php">STATS</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/Payment/index.php">PAYMENT</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="/Botnet/index.php">BOTNET</a>
                        </li>
                    </ul>
                </div>
            </nav>
        </div>

        <div class="container">
            <div class="row">
                <div class="sidenav" style="margin-top: 86px; width: auto; background-color: #242640">
                    <a href="index.php">Active Victims</a><!-- index page -->
                    <a href="statistics.php">Statics</a> <!-- it will produced using payment rate, amount of collected money... -->
                </div>
                <div class="col-sm">
                    HEY MAP MAP
                    <?php print(createVictimTable());?>
                </div>
            </div>
        </div>



        <script src="/js/jquery-3.5.1.slim.min.js"></script>
        <script src="/js/bootstrap.bundle.min.js"></script>
    </body>
</html>