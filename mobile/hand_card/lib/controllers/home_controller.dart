import 'dart:convert';

import 'package:geolocator/geolocator.dart';

import '../model/card.dart';
import '../service/user_secure_storage.dart';
import 'package:http/http.dart' as http;

class HomeController{

  Future<List<DiscountCard>> getCards() async {
    final jwt = await UserSecureStorage.getJwt();
    Position position = await _determinePosition();
    final url = 'http://10.0.2.2:8080/api/cards/?lat=${position.latitude}&lon=${position.longitude}';
    print(url);
    final response = await http.get(
      Uri.parse(url),
      headers: {
        'Content-Type': 'application/json; charset=UTF-8',
        'Authorization': 'Bearer $jwt',
      },
    );
    
    final body = json.decode(response.body);
    return body.map<DiscountCard>(DiscountCard.fromJson).toList();
  }

  Future<Position> _determinePosition() async {
    LocationPermission permission;

    permission = await Geolocator.checkPermission();

    if(permission == LocationPermission.denied) {
      permission = await Geolocator.requestPermission();
      if(permission == LocationPermission.denied) {
        return Future.error('Location Permissions are denied');
      }
    }
    return await Geolocator.getCurrentPosition();
  }

}