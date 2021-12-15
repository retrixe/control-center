import 'dart:async';

import 'package:flutter/material.dart';
import 'package:controlcenter/page.dart';
import 'package:dbus/dbus.dart';

class LenovoSettingsPage extends StatelessWidget {
  const LenovoSettingsPage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  Widget build(BuildContext context) {
    return SettingsPage(
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.start,
          children: const <Widget>[
            LenovoConservationModeSetting(),
          ],
        ),
      ),
      title: title,
    );
  }
}

class LenovoConservationModeSetting extends StatefulWidget {
  const LenovoConservationModeSetting({Key? key}) : super(key: key);

  @override
  createState() => _LenovoConservationModeSettingState();
}

class _LenovoConservationModeSettingState
    extends State<LenovoConservationModeSetting> {
  bool enabled = false;
  bool available = false;
  late Timer _timer;

  Future<void> updateStateFromDBus() async {
    var client = DBusClient.system();
    var object = DBusRemoteObject(client,
        name: 'com.retrixe.ControlCenter.v0',
        path: DBusObjectPath('/com/retrixe/ControlCenter/v0'));
    var result = await object.callMethod(
        'com.retrixe.ControlCenter.v0', 'LenovoGetConservationModeStatus', []);
    num value = result.values[0].toNative();
    setState(() {
      enabled = value == 1;
      available = value == 0 || value == 1;
    });
  }

  Future<bool> setConservationMode(bool value) async {
    var client = DBusClient.system();
    var object = DBusRemoteObject(client,
        name: 'com.retrixe.ControlCenter.v0',
        path: DBusObjectPath('/com/retrixe/ControlCenter/v0'));
    var result = await object.callMethod('com.retrixe.ControlCenter.v0',
        'LenovoSetConservationMode', [DBusBoolean(value)]);
    return result.values[0].toNative();
  }

  @override
  void initState() {
    // TODO: Error handling.
    super.initState();
    updateStateFromDBus().catchError(print);
    _timer = Timer.periodic(const Duration(seconds: 5), (timer) {
      updateStateFromDBus().catchError(print);
    });
  }

  @override
  void dispose() {
    super.dispose();
    _timer.cancel();
  }

  void toggleHandler(bool value) {
    setState(() {
      enabled = value;
    });
    setConservationMode(value).then((success) {
      if (!success) {
        // Revert the set.
        setState(() {
          enabled = !value;
        });
      }
    }).catchError(print);
  }

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Text('Conservation Mode', style: Theme.of(context).textTheme.subtitle1),
        Switch(
          value: enabled,
          onChanged: available ? toggleHandler : null,
        )
      ],
    );
  }
}
