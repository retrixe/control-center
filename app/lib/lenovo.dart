import 'dart:async';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:control_center/page.dart';
import 'package:dbus/dbus.dart';

class LenovoSettingsPage extends SettingsPage {
  const LenovoSettingsPage(
      {Key? key, required String title, required DBusClient client})
      : super(key: key, title: title, client: client);

  @override
  Widget buildPage(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          LenovoConservationModeSetting(
              client: client, showDBusError: () => showDBusError(context)),
        ],
      ),
    );
  }
}

class LenovoConservationModeSetting extends StatefulWidget {
  const LenovoConservationModeSetting(
      {Key? key, required this.client, required this.showDBusError})
      : super(key: key);

  final DBusClient client;
  final Future<void> Function() showDBusError;

  @override
  createState() => _LenovoConservationModeSettingState();
}

class _LenovoConservationModeSettingState
    extends State<LenovoConservationModeSetting> {
  bool enabled = false;
  bool available = false;
  late Timer _timer;

  Future<void> updateStateFromDBus() async {
    var object = DBusRemoteObject(widget.client,
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
    var object = DBusRemoteObject(widget.client,
        name: 'com.retrixe.ControlCenter.v0',
        path: DBusObjectPath('/com/retrixe/ControlCenter/v0'));
    var result = await object.callMethod('com.retrixe.ControlCenter.v0',
        'LenovoSetConservationMode', [DBusBoolean(value)]);
    return result.values[0].toNative();
  }

  @override
  void initState() {
    super.initState();
    updateStateFromDBus().catchError((error) {
      stderr.writeln(error);
      widget.showDBusError();
    });
    _timer = Timer.periodic(const Duration(seconds: 5), (timer) async {
      try {
        await updateStateFromDBus();
      } catch (error) {
        stderr.writeln(error);
        widget.showDBusError();
      }
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
        setState(() {
          enabled = !value;
        });
      }
    }).catchError((error) {
      setState(() {
        enabled = !value;
      });
      stderr.writeln(error);
      widget.showDBusError();
    });
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
