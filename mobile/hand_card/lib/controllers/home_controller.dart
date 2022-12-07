import 'dart:convert';

import '../model/card.dart';
import '../service/user_secure_storage.dart';
import 'package:http/http.dart' as http;

class HomeController{

  Future<List<DiscountCard>> getCards() async {
    final jwt = await UserSecureStorage.getJwt();
    final url = 'http://10.0.2.2:8080/api/cards/?lat=82.897918&lon=54.980332';
    final response = await http.get(
      Uri.parse(url),
      headers: {
        'Content-Type': 'application/json; charset=UTF-8',
        'Authorization': 'Bearer $jwt',
      },
    );
    
    final body = json.decode(response.body);
    print(body);
    return body.map<DiscountCard>(DiscountCard.fromJson).toList();
  }

    Future deleteCard(int idCard) async {
    final jwt = await UserSecureStorage.getJwt();
    final url = 'http://10.0.2.2:8080/api/cards/1';
    final response = await http.delete(
      Uri.parse(url),
      headers: {
        'Content-Type': 'application/json; charset=UTF-8',
        'Authorization': 'Bearer $jwt',
      },
    );
    final body = json.decode(response.body);
    return body.map<DiscountCard>(DiscountCard.fromJson).toList();
  }

}