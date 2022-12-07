
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class UserSecureStorage {
  static final storage = FlutterSecureStorage();
  static const keyJwt = 'jwt';

static Future setJwt(String jwt) async =>
  await storage.write(key: keyJwt, value: jwt);

static Future<String?> getJwt() async =>
  await storage.read(key: keyJwt);
}